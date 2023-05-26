package routes

import (
	"context"
	"net/http"

	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/routes/dto"

	"github.com/gin-gonic/gin"
)

func UpdateItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	body := dto.UpdateItemRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateItem(context.Background(), &pb.UpdateItemRequest{
		Id:       body.Id,
		Name:     body.Name,
		Quantity: body.Quantity,
		Price:    body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
