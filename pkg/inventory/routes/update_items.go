package routes

import (
	"context"
	"net/http"

	"api-gateway/pkg/inventory/pb"

	"github.com/gin-gonic/gin"
)

type UpdateItemRequestBody struct {
	Id       int64  `json:"product_id"`
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
}

func UpdateItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	body := UpdateItemRequestBody{}

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
