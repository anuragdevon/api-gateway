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

	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/pb/mocks"
)

func TestUpdateItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockClient := new(mocks.InventoryServiceClient)

	var inventoryServiceClient pb.InventoryServiceClient = mockClient
	router.PUT("/inventory/:id", func(ctx *gin.Context) {
		UpdateItem(ctx, inventoryServiceClient)
	})

	t.Run("UpdateItem method to return status 200 OK for successful update of an item in the inventory", func(t *testing.T) {
		itemID := "123"
		requestBody := UpdateItemRequestBody{
			Id:       123,
			Name:     "Updated Product",
			Quantity: 20,
			Price:    200,
		}
		jsonBody := `{"product_id":123,"name":"Updated Product","quantity":20,"price":200}`

		req, err := http.NewRequest("PUT", "/inventory/"+itemID, strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.UpdateItemRequest{
			Id:       requestBody.Id,
			Name:     requestBody.Name,
			Quantity: requestBody.Quantity,
			Price:    requestBody.Price,
		}
		expectedResponse := &pb.UpdateItemResponse{
			Status: 200,
		}
		mockClient.On("UpdateItem", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "UpdateItem", mock.Anything, expectedRequest)
	})

	t.Run("UpdateItem method to return status 400 BadRequest for invalid request", func(t *testing.T) {
		itemID := "456"

		jsonBody := `{"name":"Test Product","quantity":"10","price":"100"}`
		req, err := http.NewRequest("PUT", "/inventory/"+itemID, strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "UpdateItem")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("UpdateItem method to return status 502 BadGateway for an error from the inventory service", func(t *testing.T) {

		itemID := "789"
		requestBody := UpdateItemRequestBody{
			Id:       789,
			Name:     "Updated Product",
			Quantity: 20,
			Price:    200,
		}
		jsonBody := `{"product_id":789,"name":"Updated Product","quantity":20,"price":200}`

		req, err := http.NewRequest("PUT", "/inventory/"+itemID, strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.UpdateItemRequest{
			Id:       requestBody.Id,
			Name:     requestBody.Name,
			Quantity: requestBody.Quantity,
			Price:    requestBody.Price,
		}
		mockClient.On("UpdateItem", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		router.ServeHTTP(recorder, req)

		mockClient.AssertCalled(t, "UpdateItem", mock.Anything, expectedRequest)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})
}
