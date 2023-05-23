package routes

import (
	"context"
	"net/http"
	"strconv"

	"api-gateway/pkg/inventory/pb"

	"github.com/gin-gonic/gin"
)

func DeleteItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.DeleteItem(context.Background(), &pb.DeleteItemRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
