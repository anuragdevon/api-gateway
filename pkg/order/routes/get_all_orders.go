package routes

import (
	"context"
	"errors"
	"net/http"

	"api-gateway/pkg/order/pb"

	authpb "api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(ctx *gin.Context, c pb.OrderServiceClient) {
	userType, _ := ctx.Get("UserType")
	if userType != authpb.UserType_CUSTOMER {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}

	userId := ctx.GetInt64("UserId")

	res, err := c.GetAllOrders(context.Background(), &pb.GetAllOrdersRequest{
		UserId: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
