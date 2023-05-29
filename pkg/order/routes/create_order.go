package routes

import (
	"context"
	"errors"
	"net/http"

	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/routes/dto"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	if ctx.GetString("UserType") != "CUSTOMER" {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}

	userId := ctx.GetInt64("UserId")
	body := dto.CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		UserId:   userId,
		ItemId:   body.ItemId,
		Quantity: body.Quantity,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
