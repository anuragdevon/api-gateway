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

func TestDeleteItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("DeleteItem method to return status 200 OK for successfully deletion of item", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.InventoryServiceClient)

		var inventoryServiceClient pb.InventoryServiceClient = mockClient
		router.DELETE("/inventory/:id", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_ADMIN)
			DeleteItem(ctx, inventoryServiceClient)
		})

		itemID := "123"
		req, err := http.NewRequest("DELETE", "/inventory/"+itemID, nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.DeleteItemRequest{
			Id: 123,
		}
		expectedResponse := &pb.DeleteItemResponse{
			Status: 200,
		}
		mockClient.On("DeleteItem", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "DeleteItem", mock.Anything, expectedRequest)
	})

	t.Run("DeleteItem method to return status 502 BadGateway for an error from the inventory service", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.InventoryServiceClient)

		var inventoryServiceClient pb.InventoryServiceClient = mockClient
		router.DELETE("/inventory/:id", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_ADMIN)
			DeleteItem(ctx, inventoryServiceClient)
		})

		itemID := "456"
		req, err := http.NewRequest("DELETE", "/inventory/"+itemID, nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.DeleteItemRequest{
			Id: 456,
		}
		mockClient.On("DeleteItem", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		router.ServeHTTP(recorder, req)

		mockClient.AssertCalled(t, "DeleteItem", mock.Anything, expectedRequest)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})

	t.Run("DeleteItem method to return status 403 StatusForbidden for Non-Admin user type", func(t *testing.T) {
		router := gin.New()

		mockClient := new(mocks.InventoryServiceClient)

		var inventoryServiceClient pb.InventoryServiceClient = mockClient
		router.DELETE("/inventory/:id", func(ctx *gin.Context) {
			ctx.Set("UserType", authpb.UserType_CUSTOMER)
			DeleteItem(ctx, inventoryServiceClient)
		})

		itemID := "456"
		req, err := http.NewRequest("DELETE", "/inventory/"+itemID, nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.DeleteItemRequest{
			Id: 456,
		}
		mockClient.On("DeleteItem", mock.Anything, expectedRequest).Return(nil, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}
