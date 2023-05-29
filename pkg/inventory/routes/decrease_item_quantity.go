package routes

import (
	"context"
	"errors"
	"net/http"

	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/routes/dto"

	"github.com/gin-gonic/gin"
)

func DecreaseItemQuantity(ctx *gin.Context, c pb.InventoryServiceClient) {
	if ctx.GetString("UserType") != "ADMIN" {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}
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
