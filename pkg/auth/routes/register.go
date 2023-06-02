package routes

import (
	"context"
	"errors"
	"net/http"

	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/auth/routes/dto"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := dto.RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userType, ok := dto.UserTypeMap[body.UserType]
	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid user type"))
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
		UserType: userType,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
