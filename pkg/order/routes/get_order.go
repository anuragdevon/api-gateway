package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"api-gateway/pkg/order/pb"

	authpb "api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

func GetOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	userType, _ := ctx.Get("UserType")
	if userType != authpb.UserType_CUSTOMER {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}

	userId := ctx.GetInt64("UserId")
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.GetOrder(context.Background(), &pb.GetOrderRequest{
		UserId: userId,
		Id:     id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
