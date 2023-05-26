package routes

import (
	"context"
	"net/http"

	"api-gateway/pkg/order/pb"
	"api-gateway/pkg/order/routes/dto"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := dto.CreateOrderRequestBody{}

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
