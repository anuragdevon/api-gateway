package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	authpb "api-gateway/pkg/auth/pb"
	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/pb/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAllOrders method should return status 200 OK and order data for a valid user ID", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.GET("/orders", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(123))
			GetAllOrders(ctx, orderServiceClient)
		})

		userID := "123"
		req, err := http.NewRequest("GET", "/orders?user_id="+userID, nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetAllOrdersRequest{
			UserId: 123,
		}
		expectedResponse := &pb.GetAllOrdersResponse{
			Status: 200,
			Data: []*pb.GetAllOrdersData{
				{
					Id:       1,
					ItemId:   100,
					Quantity: 2,
				},
				{
					Id:       2,
					ItemId:   200,
					Quantity: 3,
				},
			},
		}
		mockClient.On("GetAllOrders", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200,"data":[{"id":1,"itemId":100,"quantity":2},{"id":2,"itemId":200,"quantity":3}]}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "GetAllOrders", mock.Anything, expectedRequest)
	})

	t.Run("GetAllOrders method should return status 502 Bad Gateway for an error from the order service", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.OrderServiceClient)

		var orderServiceClient pb.OrderServiceClient = mockClient
		router.GET("/orders", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(123))
			GetAllOrders(ctx, orderServiceClient)
		})

		userID := "456"
		req, err := http.NewRequest("GET", "/orders?user_id="+userID, nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetAllOrdersRequest{
			UserId: 123,
		}
		mockClient.On("GetAllOrders", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		router.ServeHTTP(recorder, req)

		mockClient.AssertCalled(t, "GetAllOrders", mock.Anything, expectedRequest)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})
}
