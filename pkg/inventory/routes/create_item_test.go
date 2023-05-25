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

func TestCreateItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockClient := new(mocks.InventoryServiceClient)

	var inventoryServiceClient pb.InventoryServiceClient = mockClient
	router.POST("/inventory", func(ctx *gin.Context) {
		CreateItem(ctx, inventoryServiceClient)
	})

	t.Run("CreateItem Method to return status 201 StatusCreated for successful item creation in inventory", func(t *testing.T) {

		requestBody := dto.CreateItemRequestBody{
			Name:     "Test Product",
			Quantity: 10,
			Price:    100,
		}

		expectedRequest := &pb.CreateItemRequest{
			Name:     requestBody.Name,
			Quantity: requestBody.Quantity,
			Price:    requestBody.Price,
		}

		expectedResponse := &pb.CreateItemResponse{
			Id:     1,
			Status: 201,
		}

		mockClient.On("CreateItem", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		jsonBody := `{"name":"Test Product","quantity":10,"price":100}`

		req, err := http.NewRequest("POST", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)
		// TODO: id check
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("CreateItem Method to return status 400 BadRequest for invalid request", func(t *testing.T) {
		jsonBody := `{"name":"Test Product","quantity":"10","price":"100"}`
		req, err := http.NewRequest("POST", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "CreateItem")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("CreateItem Method to return status 502 BadGateway for bad gateway error", func(t *testing.T) {
		expectedRequest := &pb.CreateItemRequest{
			Name:     "Test Product",
			Quantity: 10,
			Price:    100,
		}

		mockClient.On("CreateItem", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		jsonBody := `{"name":"Test Product","quantity":10,"price":100}`
		req, err := http.NewRequest("POST", "/inventory", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})
}
