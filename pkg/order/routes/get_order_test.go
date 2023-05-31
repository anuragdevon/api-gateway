package routes

import (
	authpb "api-gateway/pkg/auth/pb"
	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/pb/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetOrder method should return status 200 OK and order data for a valid order ID", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.GET("/order/:id", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(123))
			GetOrder(ctx, orderServiceClient)
		})

		orderID := int64(123)
		req, err := http.NewRequest("GET", "/order/123", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetOrderRequest{
			UserId: 123,
			Id:     orderID,
		}
		expectedResponse := &pb.GetOrderResponse{
			Status: 200,
			Data: &pb.GetOrderData{
				Id:       orderID,
				ItemId:   123,
				Name:     "Test Product",
				Quantity: 10,
				Price:    100,
			},
		}
		mockClient.On("GetOrder", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200,"data":{"id":123,"itemId":123,"name":"Test Product","quantity":10,"price":100}}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "GetOrder", mock.Anything, expectedRequest)
	})

	t.Run("GetOrder method should return status 502 Bad Gateway for an error from the order service", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.GET("/order/:id", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(456))
			GetOrder(ctx, orderServiceClient)
		})

		orderID := int64(456)
		req, err := http.NewRequest("GET", "/order/456", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetOrderRequest{
			UserId: 456,
			Id:     orderID,
		}
		mockClient.On("GetOrder", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		router.ServeHTTP(recorder, req)

		mockClient.AssertCalled(t, "GetOrder", mock.Anything, expectedRequest)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})
}
