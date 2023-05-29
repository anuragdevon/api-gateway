package routes

import (
	"context"
	"net/http"

	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/routes/dto"

	"github.com/gin-gonic/gin"
)

func DecreaseItemQuantity(ctx *gin.Context, c pb.InventoryServiceClient) {
	body := dto.DecreaseItemQuantityRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.DecreaseItemQuantity(context.Background(), &pb.DecreaseItemQuantityRequest{
		Id:       body.Id,
		Quantity: body.Quantity,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
