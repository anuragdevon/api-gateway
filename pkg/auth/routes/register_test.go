package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/auth/routes/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AuthServiceClient interface {
	pb.AuthServiceClient
}

func TestRegister(t *testing.T) {
	router := gin.Default()
	t.Run("Register Method to return status 200 StatusOK for successful registration", func(t *testing.T) {
		mockClient := new(mocks.MockAuthServiceClient)

		requestBody := RegisterRequestBody{
			Email:    "test@example.com",
			Password: "password",
		}

		expectedRequest := &pb.RegisterRequest{
			Email:    requestBody.Email,
			Password: requestBody.Password,
		}

		expectedResponse := &pb.RegisterResponse{
			Status: 200,
		}

		mockClient.On("Register", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		jsonBody := `{"email":"test@example.com","password":"password"}`
		req, err := http.NewRequest("POST", "/register", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		var authServiceClient AuthServiceClient = mockClient

		router.POST("/register", func(ctx *gin.Context) {
			Register(ctx, authServiceClient)
		})

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())
	})

	t.Run("Register Method to return status 400 BadRequest for invalid request body", func(t *testing.T) {
		mockClient := new(mocks.MockAuthServiceClient)

		jsonBody := `{}`

		req, err := http.NewRequest("POST", "/register", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		var authServiceClient AuthServiceClient = mockClient

		router := gin.Default()

		router.POST("/register", func(ctx *gin.Context) {
			Register(ctx, authServiceClient)
		})

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
