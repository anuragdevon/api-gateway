package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	authpb "api-gateway/pkg/auth/pb"
	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/pb/mocks"
)

func TestGetAllItems(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAllItems method to return status 200 OK and all item details", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.InventoryServiceClient)

		var inventoryServiceClient pb.InventoryServiceClient = mockClient
		router.GET("/inventory", func(ctx *gin.Context) {
			ctx.Set("UserId", int64(123))
			GetAllItems(ctx, inventoryServiceClient)
		})

		req, err := http.NewRequest("GET", "/inventory", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetAllItemsRequest{
			Userid: 123,
		}
		expectedResponse := &pb.GetAllItemsResponse{
			Status: 200,
			Data: []*pb.GetItemData{
				{
					Id:       1,
					Name:     "Item 1",
					Quantity: 10,
					Price:    100,
				},
				{
					Id:       2,
					Name:     "Item 2",
					Quantity: 20,
					Price:    200,
				},
			},
		}
		mockClient.On("GetAllItems", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200,"data":[{"id":1,"name":"Item 1","quantity":10,"price":100},{"id":2,"name":"Item 2","quantity":20,"price":200}]}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "GetAllItems", mock.Anything, expectedRequest)
	})

	t.Run("GetAllItems method to return status 502 BadGateway for an error from the inventory service", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.InventoryServiceClient)

		var inventoryServiceClient pb.InventoryServiceClient = mockClient
		router.GET("/inventory", func(ctx *gin.Context) {
			ctx.Set("UserId", int64(123))
			GetAllItems(ctx, inventoryServiceClient)
		})

		req, err := http.NewRequest("GET", "/inventory", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetAllItemsRequest{
			Userid: 123,
		}
		mockClient.On("GetAllItems", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		router.ServeHTTP(recorder, req)

		mockClient.AssertCalled(t, "GetAllItems", mock.Anything, expectedRequest)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})

	t.Run("GetAllItems method to return status 403 StatusForbidden for Non-Admin user type", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.InventoryServiceClient)

		var inventoryServiceClient pb.InventoryServiceClient = mockClient
		router.GET("/inventory", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			ctx.Set("UserId", int64(123))
			GetAllItems(ctx, inventoryServiceClient)
		})

		req, err := http.NewRequest("GET", "/inventory", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}
