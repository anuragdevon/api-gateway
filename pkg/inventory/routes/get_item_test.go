package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/pb/mocks"
)

func TestGetItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockClient := new(mocks.InventoryServiceClient)

	var inventoryServiceClient pb.InventoryServiceClient = mockClient
	router.GET("/inventory/:id", func(ctx *gin.Context) {
		GetItem(ctx, inventoryServiceClient)
	})

	t.Run("GetItem method to return status 200 OK and item details for a valid item ID", func(t *testing.T) {
		itemID := "123"
		req, err := http.NewRequest("GET", "/inventory/"+itemID, nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetItemRequest{
			Id: 123,
		}
		expectedResponse := &pb.GetItemResponse{
			Status: 200,
			Data: &pb.GetItemData{
				Id:       123,
				Name:     "Test Product",
				Quantity: 10,
				Price:    100,
			},
		}
		mockClient.On("GetItem", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200,"data":{"id":123,"name":"Test Product","quantity":10,"price":100}}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "GetItem", mock.Anything, expectedRequest)
	})

	t.Run("GetItem method to return status 502 BadGateway for an error from the inventory service", func(t *testing.T) {

		itemID := "456"
		req, err := http.NewRequest("GET", "/inventory/"+itemID, nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.GetItemRequest{
			Id: 456,
		}
		mockClient.On("GetItem", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		router.ServeHTTP(recorder, req)

		mockClient.AssertCalled(t, "GetItem", mock.Anything, expectedRequest)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)

	})
}
