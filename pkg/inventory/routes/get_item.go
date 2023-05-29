package routes

import (
	"context"
	"net/http"
	"strconv"

	"api-gateway/pkg/inventory/pb"

	"github.com/gin-gonic/gin"
)

func GetItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.GetItem(context.Background(), &pb.GetItemRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
