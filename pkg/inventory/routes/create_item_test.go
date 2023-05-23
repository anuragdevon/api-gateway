package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/routes/mocks"
)

func TestCreateItem(t *testing.T) {
	router := gin.Default()
	t.Run("CreateItem method to return status 201 Created for successful entry of item in inventory", func(t *testing.T) {

		mockClient := new(mocks.InventoryServiceClient)

		router.POST("/inventory", func(ctx *gin.Context) {
			CreateItem(ctx, mockClient)
		})

		requestBody := CreateItemRequestBody{
			Name:     "Test Product",
			Quantity: 10,
			Price:    100,
		}

		jsonBody := `{"name":"Test Product","quantity":10,"price":100}`

		req, err := http.NewRequest("POST", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		expectedRequest := &pb.CreateItemRequest{
			Name:     requestBody.Name,
			Quantity: requestBody.Quantity,
			Price:    requestBody.Price,
		}
		expectedResponse := &pb.CreateItemResponse{
			Status: 201,
		}
		mockClient.On("CreateItem", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		expectedResponseBody := `{"status":201}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())

		mockClient.AssertCalled(t, "CreateItem", mock.Anything, expectedRequest)
	})

	t.Run("CreateItem Method to return status 400 BadRequest for invalid request", func(t *testing.T) {
		mockClient := new(mocks.InventoryServiceClient)

		jsonBody := `{"name":"new-test-product"}`
		req, err := http.NewRequest("POST", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		var inventoryServiceClient mocks.InventoryServiceClient = *mockClient

		router.POST("/inventory", func(ctx *gin.Context) {
			CreateItem(ctx, &inventoryServiceClient)
		})

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "CreateItem")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		expectedResponseBody := `{"error":"invalid request body"}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())
	})
}
