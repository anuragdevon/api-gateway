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
	"api-gateway/pkg/inventory/routes/dto"
)

func TestDecreaseItemQuantity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockClient := new(mocks.InventoryServiceClient)

	var inventoryServiceClient pb.InventoryServiceClient = mockClient
	router.PUT("/inventory", func(ctx *gin.Context) {
		DecreaseItemQuantity(ctx, inventoryServiceClient)
	})

	t.Run("DecreaseItemQuantity method to return status 200 OK for successful update of an item in the inventory", func(t *testing.T) {
		requestBody := dto.DecreaseItemQuantityRequestBody{
			Id:       123,
			Quantity: 10,
		}
		jsonBody := `{"product_id":123,"quantity":10}`

		req, err := http.NewRequest("PUT", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.DecreaseItemQuantityRequest{
			Id:       requestBody.Id,
			Quantity: requestBody.Quantity,
		}
		expectedResponse := &pb.DecreaseItemQuantityResponse{
			Status: 200,
		}
		mockClient.On("DecreaseItemQuantity", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "DecreaseItemQuantity", mock.Anything, expectedRequest)
	})

	t.Run("DecreaseItemQuantity method to return status 400 BadRequest for invalid request", func(t *testing.T) {
		jsonBody := `{"product_id":"1","quantity":"10"}`
		req, err := http.NewRequest("PUT", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "DecreaseItemQuantity")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("DecreaseItemQuantity method to return status 502 BadGateway for an error from the inventory service", func(t *testing.T) {
		requestBody := dto.DecreaseItemQuantityRequestBody{
			Id:       789,
			Quantity: 20,
		}
		jsonBody := `{"product_id":789,"quantity":20}`

		req, err := http.NewRequest("PUT", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.DecreaseItemQuantityRequest{
			Id:       requestBody.Id,
			Quantity: requestBody.Quantity,
		}
		mockClient.On("DecreaseItemQuantity", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		router.ServeHTTP(recorder, req)

		mockClient.AssertCalled(t, "DecreaseItemQuantity", mock.Anything, expectedRequest)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})
}
