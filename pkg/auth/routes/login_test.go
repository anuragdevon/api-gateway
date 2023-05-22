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

func TestLogin(t *testing.T) {
	router := gin.Default()
	t.Run("Login Method to return status 201 StatusCreated for successful login", func(t *testing.T) {
		mockClient := new(mocks.MockAuthServiceClient)

		requestBody := LoginRequestBody{
			Email:    "test@example.com",
			Password: "password",
		}

		expectedRequest := &pb.LoginRequest{
			Email:    requestBody.Email,
			Password: requestBody.Password,
		}

		expectedResponse := &pb.LoginResponse{
			Token: "some_expected_valid_token",
		}

		mockClient.On("Login", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		jsonBody := `{"email":"test@example.com","password":"password"}`
		req, err := http.NewRequest("POST", "/login", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		var authServiceClient AuthServiceClient = mockClient

		router.POST("/login", func(ctx *gin.Context) {
			Login(ctx, authServiceClient)
		})

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		expectedResponseBody := `{"token":"some_expected_valid_token"}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())
	})

	t.Run("Login Method to return status 400 BadRequest for invalid request", func(t *testing.T) {
		mockClient := new(mocks.MockAuthServiceClient)

		jsonBody := `{"email":"test@example.com"}`
		req, err := http.NewRequest("POST", "/login", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		var authServiceClient AuthServiceClient = mockClient

		router.POST("/login", func(ctx *gin.Context) {
			Login(ctx, authServiceClient)
		})

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "Login")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		expectedResponseBody := `{"error":"invalid request body"}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())
	})
}
