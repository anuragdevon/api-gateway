package routes

import (
	"context"
	"net/http"

	"api-gateway/pkg/inventory/pb"

	"github.com/gin-gonic/gin"
)

func GetAllItems(ctx *gin.Context, c pb.InventoryServiceClient) {

	res, err := c.GetAllItems(context.Background(), &pb.GetAllItemsRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
