package inventory

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/inventory/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/inventory")
	routes.POST("", svc.CreateItem)
	routes.GET("/:id", svc.GetItem)
	routes.GET("", svc.GetAllItems)
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
