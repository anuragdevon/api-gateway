package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"api-gateway/pkg/auth/routes/mocks"
)

func TestAuthRequired(t *testing.T) {
	router := gin.Default()

	t.Run("AuthRequired middleware should return status 401 Unauthorized when authorization header is missing", func(t *testing.T) {
		mockClient := new(mocks.MockAuthServiceClient)

		authMiddleware := InitAuthMiddleware(&ServiceClient{Client: mockClient})

		router.Use(authMiddleware.AuthRequired)

		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})
}
