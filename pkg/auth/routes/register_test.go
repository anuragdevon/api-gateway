package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	pb "api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"api-gateway/pkg/auth/pb/mocks"
)

func TestRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockAuthServiceClient(ctrl)

	router := gin.Default()

	Register(router.Group("/api"), mockClient)

	requestBody := RegisterRequestBody{
		Email:    "test@example.com",
		Password: "password",
		UserType: pb.UserType_CUSTOMER,
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/register", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("json", requestBody)

	mockClient.EXPECT().Register(gomock.Any(), gomock.Any()).Return(&pb.RegisterResponse{
		Status: 200,
	}, nil)

	router.HandleContext(ctx)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"status":200}`
	assert.Equal(t, expectedResponse, w.Body.String())
}
