package routes

import (
	"api-gateway/pkg/inventory/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllInventoryItems(ctx *gin.Context, c pb.InventoryServiceClient) {
	res, err := c.GetAllInventoryItems(context.Background(), &pb.GetAllInventoryItemsRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
