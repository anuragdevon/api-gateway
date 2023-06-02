package routes

import (
	"api-gateway/pkg/inventory/pb"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	authpb "api-gateway/pkg/auth/pb"
)

func GetAllItems(ctx *gin.Context, c pb.InventoryServiceClient) {
	userType, _ := ctx.Get("UserType")
	if userType != authpb.UserType_ADMIN {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}

	userId := ctx.GetInt64("UserId")

	res, err := c.GetAllItems(context.Background(), &pb.GetAllItemsRequest{
		Userid: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
