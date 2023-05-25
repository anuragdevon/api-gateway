package routes

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"

// 	"api-gateway/pkg/inventory/pb"
// 	"api-gateway/pkg/inventory/routes/mocks"
// )

// func TestGetItem(t *testing.T) {
// 	router := gin.Default()
// 	t.Run("GetItem method to return status 200 OK and item details for a valid item ID", func(t *testing.T) {
// 		mockClient := new(mocks.InventoryServiceClient)

// 		router.GET("/inventory/:id", func(ctx *gin.Context) {
// 			GetItem(ctx, mockClient)
// 		})

// 		itemID := "123"
// 		req, err := http.NewRequest("GET", "/inventory/"+itemID, nil)
// 		assert.NoError(t, err)

// 		recorder := httptest.NewRecorder()

// 		expectedRequest := &pb.GetItemRequest{
// 			Id: 123,
// 		}
// 		expectedResponse := &pb.GetItemResponse{
// 			Status: 200,
// 			Data: &pb.GetItemData{
// 				Id:       123,
// 				Name:     "Test Product",
// 				Quantity: 10,
// 				Price:    100,
// 			},
// 		}
// 		mockClient.On("GetItem", mock.Anything, expectedRequest).Return(expectedResponse, nil)

// 		router.ServeHTTP(recorder, req)

// 		assert.Equal(t, http.StatusOK, recorder.Code)

// 		expectedResponseBody := `{"status":200,"data":{"id":123,"name":"Test Product","quantity":10,"price":100}}`
// 		assert.Equal(t, expectedResponseBody, recorder.Body.String())

// 		mockClient.AssertCalled(t, "GetItem", mock.Anything, expectedRequest)
// 	})

// 	t.Run("GetItem method to return status 500 InternalServerError for an error from the inventory service", func(t *testing.T) {
// 		mockClient := new(mocks.InventoryServiceClient)

// 		itemID := "456"
// 		req, err := http.NewRequest("GET", "/inventory/"+itemID, nil)
// 		assert.NoError(t, err)

// 		recorder := httptest.NewRecorder()

// 		var inventoryServiceClient mocks.InventoryServiceClient = *mockClient

// 		router.GET("/inventory/:id", func(ctx *gin.Context) {
// 			GetItem(ctx, &inventoryServiceClient)
// 		})

// 		expectedRequest := &pb.GetItemRequest{
// 			Id: 456,
// 		}
// 		mockClient.On("GetItem", mock.Anything, expectedRequest).Return(nil, "someError")

// 		router.ServeHTTP(recorder, req)

// 		mockClient.AssertCalled(t, "GetItem", mock.Anything, expectedRequest)

// 		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

// 	})
// }
