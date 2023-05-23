package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/routes/mocks"
)

func TestCreateOrder(t *testing.T) {
	router := gin.Default()

	t.Run("CreateOrder method should return status 201 Created for a successful order creation", func(t *testing.T) {
		mockClient := new(mocks.OrderServiceClient)

		router.POST("/order", func(ctx *gin.Context) {
			CreateOrder(ctx, mockClient)
		})

		requestBody := CreateOrderRequestBody{
			UserId:   12,
			ItemId:   123,
			Quantity: 10,
		}

		jsonBody := `{"itemId": 123, "quantity": 10, "user_id": 12}`

		req, err := http.NewRequest("POST", "/order", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.CreateOrderRequest{
			UserId:   requestBody.UserId,
			ItemId:   requestBody.ItemId,
			Quantity: requestBody.Quantity,
		}
		expectedResponse := &pb.CreateOrderResponse{
			Status: 201,
		}
		mockClient.On("CreateOrder", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		expectedResponseBody := `{"status":201}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "CreateOrder", mock.Anything, expectedRequest)
	})

	t.Run("CreateOrder method should return status 400 BadRequest for invalid request", func(t *testing.T) {
		mockClient := new(mocks.OrderServiceClient)

		jsonBody := `{"itemId": "abc"}`
		req, err := http.NewRequest("POST", "/orders", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		var orderServiceClient mocks.OrderServiceClient = *mockClient

		router.POST("/orders", func(ctx *gin.Context) {
			CreateOrder(ctx, &orderServiceClient)
		})

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "CreateOrder")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		expectedResponseBody := `{"error":"invalid request body"}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())
	})
}
