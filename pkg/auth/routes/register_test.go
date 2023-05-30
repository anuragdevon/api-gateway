package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/auth/routes/dto"
	"api-gateway/pkg/auth/routes/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AuthServiceClient interface {
	pb.AuthServiceClient
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockClient := new(mocks.MockAuthServiceClient)

	var authServiceClient AuthServiceClient = mockClient
	router.POST("/register", func(ctx *gin.Context) {
		Register(ctx, authServiceClient)
	})

	t.Run("Register Method to return status 200 StatusOK for successful registration", func(t *testing.T) {
		requestBody := dto.RegisterRequestBody{
			Email:    "testregister@example.com",
			Password: "password",
			UserType: "admin",
		}

		expectedRequest := &pb.RegisterRequest{
			Email:    requestBody.Email,
			Password: requestBody.Password,
			UserType: dto.UserTypeMap[requestBody.UserType],
		}

		expectedResponse := &pb.RegisterResponse{
			Status: 200,
		}

		mockClient.On("Register", mock.Anything, expectedRequest).Return(expectedResponse, nil)

		jsonBody := `{"email":"testregister@example.com","password":"password", "user_type": "admin"}`
		req, err := http.NewRequest("POST", "/register", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedResponseBody := `{"status":200}`
		assert.Equal(t, expectedResponseBody, recorder.Body.String())
	})

	t.Run("Register Method to return status 400 BadRequest for invalid request", func(t *testing.T) {
		jsonBody := `{"email": 1, "password": 2}`
		req, err := http.NewRequest("POST", "/register", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		router.ServeHTTP(recorder, req)

		mockClient.AssertNotCalled(t, "Register")

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Register Method to return status 502 BadGateway for bad gateway error", func(t *testing.T) {
		expectedRequest := &pb.RegisterRequest{
			Email:    "testregister3@example.com",
			Password: "password",
		}

		mockClient.On("Register", mock.Anything, expectedRequest).Return(nil, errors.New("bad gateway error"))

		jsonBody := `{"email":"testregister3@example.com","password":"password"}`
		req, err := http.NewRequest("POST", "/register", strings.NewReader(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		mockClient.AssertExpectations(t)

		assert.Equal(t, http.StatusBadGateway, recorder.Code)
	})
}
