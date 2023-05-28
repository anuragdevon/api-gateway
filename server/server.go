package server

import (
	"log"

	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/inventory"
	"api-gateway/pkg/order"

	"github.com/gin-gonic/gin"
)

func Run() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, &c)
	inventory.RegisterRoutes(r, &c, authSvc)
	order.RegisterRoutes(r, &c, authSvc)

	r.Run(c.Port)
}
