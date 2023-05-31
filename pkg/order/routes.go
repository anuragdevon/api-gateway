package order

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {

	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)

	routes.POST("/", svc.CreateOrder)
	routes.GET("/:id", svc.GetOrder)
	routes.GET("", svc.GetAllOrders)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}

func (svc *ServiceClient) GetOrder(ctx *gin.Context) {
	routes.GetOrder(ctx, svc.Client)
}

func (svc *ServiceClient) GetAllOrders(ctx *gin.Context) {
	// routes.GetAllOrders(ctx, svc.Client)
}
