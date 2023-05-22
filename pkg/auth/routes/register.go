package routes

import (
	"net/http"

	"api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string      `json:"email"`
	Password string      `json:"password"`
	UserType pb.UserType `json:"user_type"`
}

func Register(routerGroup *gin.RouterGroup, client pb.AuthServiceClient) {
	routerGroup.POST("/register", func(ctx *gin.Context) {
		body := RegisterRequestBody{}

		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		res, err := client.Register(ctx, &pb.RegisterRequest{
			Email:    body.Email,
			Password: body.Password,
			UserType: body.UserType,
		})

		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return
		}

		ctx.JSON(int(res.Status), &res)
	})
}
