package routes

import (
	"context"
	"net/http"

	"api-gateway/pkg/order/pb"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ItemId   int64 `json:"itemId"`
	Quantity int64 `json:"quantity"`
	UserId   int64 `json:"user_id"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{

		UserId:   body.UserId,
		ItemId:   body.ItemId,
		Quantity: body.Quantity,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
