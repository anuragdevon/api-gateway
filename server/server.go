package server

import (
	"log"

	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/inventory"

	"github.com/gin-gonic/gin"
)

func Run() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)
	inventory.RegisterRoutes(r, &c)
	r.Run(c.Port)
}
