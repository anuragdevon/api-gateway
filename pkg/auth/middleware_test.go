package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/auth/routes/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func performRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestAuthRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("AuthRequired should set userId in context and call next handler if validation succeeds", func(t *testing.T) {
		router := gin.Default()
		mockClient := new(mocks.MockAuthServiceClient)
		authMiddleware := InitAuthMiddleware(&ServiceClient{Client: mockClient})
		router.Use(authMiddleware.AuthRequired)

		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer some_token")

		mockClient.On("Validate", mock.Anything, &pb.ValidateRequest{Token: "some_token"}).
			Return(&pb.ValidateResponse{Status: http.StatusOK, UserId: 1, UserType: pb.UserType_CUSTOMER}, nil)

		router.GET("/", func(ctx *gin.Context) {
			userId, exists := ctx.Get("UserId")
			assert.True(t, exists)
			assert.Equal(t, int64(1), userId)
			ctx.Status(http.StatusOK)
		})

		w := performRequest(router, req)

		mockClient.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("AuthRequired should return status 401 Unauthorized if authorization header is missing", func(t *testing.T) {
		router := gin.Default()
		mockClient := new(mocks.MockAuthServiceClient)
		authMiddleware := InitAuthMiddleware(&ServiceClient{Client: mockClient})
		router.Use(authMiddleware.AuthRequired)

		req, _ := http.NewRequest("GET", "/", nil)
		w := performRequest(router, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("AuthRequired should return status 401 Unauthorized if token is missing in authorization header", func(t *testing.T) {
		router := gin.Default()
		mockClient := new(mocks.MockAuthServiceClient)
		authMiddleware := InitAuthMiddleware(&ServiceClient{Client: mockClient})
		router.Use(authMiddleware.AuthRequired)

		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer")
		w := performRequest(router, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("AuthRequired should return status 401 Unauthorized if validation fails", func(t *testing.T) {
		router := gin.Default()
		mockClient := new(mocks.MockAuthServiceClient)
		authMiddleware := InitAuthMiddleware(&ServiceClient{Client: mockClient})
		router.Use(authMiddleware.AuthRequired)

		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer some_token")

		mockClient.On("Validate", mock.Anything, &pb.ValidateRequest{Token: "some_token"}).
			Return(&pb.ValidateResponse{Status: http.StatusUnauthorized}, nil)

		w := performRequest(router, req)

		mockClient.AssertExpectations(t)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
