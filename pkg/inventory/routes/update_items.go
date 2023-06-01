package routes

import (
	"context"
	"errors"
	"net/http"

	authpb "api-gateway/pkg/auth/pb"
	"api-gateway/pkg/inventory/pb"
	"api-gateway/pkg/inventory/routes/dto"

	"github.com/gin-gonic/gin"
)

func UpdateItem(ctx *gin.Context, c pb.InventoryServiceClient) {
	userType, _ := ctx.Get("UserType")
	if userType != authpb.UserType_ADMIN {
		ctx.AbortWithError(http.StatusForbidden, errors.New("invalid user type"))
		return
	}

	userId := ctx.GetInt64("UserId")

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
		Userid:   userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
