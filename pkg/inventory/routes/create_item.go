package routes

import (
	"context"
	"net/http"

	"api-gateway/pkg/inventory/pb"

	"github.com/gin-gonic/gin"
)

type CreateItemRequestBody struct {
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
}

func CreateItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	body := CreateItemRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateItem(context.Background(), &pb.CreateItemRequest{
		Name:     body.Name,
		Quantity: body.Quantity,
		Price:    body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
