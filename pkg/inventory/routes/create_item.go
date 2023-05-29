package routes

import (
	"context"
	"errors"
	"net/http"

	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/routes/dto"

	"github.com/gin-gonic/gin"
)

func CreateItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	if ctx.GetString("UseType") != "ADMIN" {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}
	body := dto.CreateItemRequestBody{}

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

	ctx.JSON(int(res.Status), &res)
}
