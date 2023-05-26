package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/pb/mocks"
	"api-gateway/pkg/order/routes/dto"
)

func TestCreateOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockClient := new(mocks.OrderServiceClient)

	var orderServiceClient pb.OrderServiceClient = mockClient
	router.POST("/order", func(ctx *gin.Context) {
		CreateOrder(ctx, orderServiceClient)
	})

	t.Run("CreateOrder method should return status 201 Created for a successful order creation", func(t *testing.T) {
		requestBody := dto.CreateOrderRequestBody{
			UserId:   12,
			ItemId:   123,
			Quantity: 10,
		}

		jsonBody := `{"item_id": 123, "quantity": 10, "user_id": 12}`

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
		jsonBody := `{"item_id": "123", "quantity": "10", "user_id": "12"}`
		req, err := http.NewRequest("POST", "/order", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "CreateOrder")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("CreateOrder Method to return status 502 BadGateway for bad gateway error", func(t *testing.T) {
		expectedRequest := &pb.CreateOrderRequest{
			UserId:   12,
			ItemId:   123,
			Quantity: 10,
		}

		mockClient.On("CreateOrder", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		jsonBody := `{"item_id": 123, "quantity": 10, "user_id": 12}`

		req, err := http.NewRequest("POST", "/order", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})
}
