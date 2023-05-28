package inventory

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/inventory/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {

	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/inventory")
	routes.Use(a.AuthRequired)

	routes.POST("", svc.CreateItem)
	routes.GET("/:id", svc.GetItem)
	routes.GET("", svc.GetAllItems)
	routes.PUT("/:id", svc.UpdateItem)
	routes.DELETE("/:id", svc.DeleteItem)
}

func (svc *ServiceClient) CreateItem(ctx *gin.Context) {
	routes.CreateItem(ctx, svc.Client)
}

func (svc *ServiceClient) GetItem(ctx *gin.Context) {
	routes.GetItem(ctx, svc.Client)
}

func (svc *ServiceClient) GetAllItems(ctx *gin.Context) {
	routes.GetAllItems(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateItem(ctx *gin.Context) {
	routes.UpdateItem(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteItem(ctx *gin.Context) {
	routes.DeleteItem(ctx, svc.Client)
}
