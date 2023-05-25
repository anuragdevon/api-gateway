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

// func TestDeleteItem(t *testing.T) {
// 	router := gin.Default()
// 	t.Run("DeleteItem method to return status 200 OK for successfully deletion of item", func(t *testing.T) {
// 		mockClient := new(mocks.InventoryServiceClient)

// 		router.DELETE("/inventory/:id", func(ctx *gin.Context) {
// 			DeleteItem(ctx, mockClient)
// 		})

// 		itemID := "123"
// 		req, err := http.NewRequest("DELETE", "/inventory/"+itemID, nil)
// 		assert.NoError(t, err)

// 		recorder := httptest.NewRecorder()

// 		expectedRequest := &pb.DeleteItemRequest{
// 			Id: 123,
// 		}
// 		expectedResponse := &pb.DeleteItemResponse{
// 			Status: 200,
// 		}
// 		mockClient.On("DeleteItem", mock.Anything, expectedRequest).Return(expectedResponse, nil)

// 		router.ServeHTTP(recorder, req)

// 		assert.Equal(t, http.StatusOK, recorder.Code)

// 		expectedResponseBody := `{"status":200}`
// 		assert.Equal(t, expectedResponseBody, recorder.Body.String())

// 		mockClient.AssertCalled(t, "DeleteItem", mock.Anything, expectedRequest)
// 	})

// 	t.Run("DeleteItem method to return status 502 BadGateway for an error from the inventory service", func(t *testing.T) {
// 		mockClient := new(mocks.InventoryServiceClient)

// 		itemID := "456"
// 		req, err := http.NewRequest("DELETE", "/inventory/"+itemID, nil)
// 		assert.NoError(t, err)

// 		recorder := httptest.NewRecorder()

// 		var inventoryServiceClient mocks.InventoryServiceClient = *mockClient

// 		router.DELETE("/inventory/:id", func(ctx *gin.Context) {
// 			DeleteItem(ctx, &inventoryServiceClient)
// 		})

// 		expectedRequest := &pb.DeleteItemRequest{
// 			Id: 456,
// 		}
// 		mockClient.On("DeleteItem", mock.Anything, expectedRequest).Return(nil, "some_internal_service_error")

// 		router.ServeHTTP(recorder, req)

// 		mockClient.AssertCalled(t, "GetItem", mock.Anything, expectedRequest)

// 		assert.Equal(t, http.StatusBadGateway, recorder.Code)

// 	})
// }
