package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	authpb "api-gateway/pkg/auth/pb"
	"api-gateway/pkg/inventory/pb"

	"github.com/gin-gonic/gin"
)

func DeleteItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	userType, _ := ctx.Get("UserType")
	if userType != authpb.UserType_ADMIN {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	userId := ctx.GetInt64("UserId")

	res, err := c.DeleteItem(context.Background(), &pb.DeleteItemRequest{
		Id:     int64(id),
		Userid: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
