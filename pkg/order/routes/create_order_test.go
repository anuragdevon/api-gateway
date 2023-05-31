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

	authpb "api-gateway/pkg/auth/pb"
	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/pb/mocks"
	"api-gateway/pkg/order/routes/dto"
)

func TestCreateOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("CreateOrder method should return status 201 Created for a successful order creation", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.POST("/order", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(123))
			CreateOrder(ctx, orderServiceClient)
		})

		requestBody := dto.CreateOrderRequestBody{
			ItemId:   123,
			Quantity: 10,
		}

		jsonBody := `{"item_id": 123, "quantity": 10}`

		req, err := http.NewRequest("POST", "/order", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.CreateOrderRequest{
			ItemId:   requestBody.ItemId,
			Quantity: requestBody.Quantity,
			UserId:   123,
		}
		expectedResponse := &pb.CreateOrderResponse{
			Status: 201,
			Id:     1,
		}
		mockClient.On("CreateOrder", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		expectedResponseBody := `{"status":201,"id":1}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "CreateOrder", mock.Anything, expectedRequest)
	})

	t.Run("CreateOrder method should return status 400 BadRequest for invalid request", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.POST("/order", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(12))
			CreateOrder(ctx, orderServiceClient)
		})

		jsonBody := `{"item_id": "123", "quantity": "10"}`
		req, err := http.NewRequest("POST", "/order", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "CreateOrder")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("CreateOrder Method to return status 502 BadGateway for bad gateway error", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.POST("/order", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(123))
			CreateOrder(ctx, orderServiceClient)
		})

		expectedRequest := &pb.CreateOrderRequest{
			ItemId:   2,
			Quantity: 10,
			UserId:   123,
		}

		mockClient.On("CreateOrder", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		jsonBody := `{"item_id": 2, "quantity": 10}`

		req, err := http.NewRequest("POST", "/order", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})

	t.Run("CreateOrder Method to return status 403 StatusForbidden for Non-Customer user type", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.POST("/order", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_ADMIN)
			ctx.Set("UserId", int64(123))
			CreateOrder(ctx, orderServiceClient)
		})

		expectedRequest := &pb.CreateOrderRequest{
			ItemId:   2,
			Quantity: 10,
			UserId:   123,
		}

		mockClient.On("CreateOrder", mock.Anything, expectedRequest).Return(nil, nil)

		jsonBody := `{"item_id": 2, "quantity": 10}`

		req, err := http.NewRequest("POST", "/order", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}
