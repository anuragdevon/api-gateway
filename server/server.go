package server

import (
	"log"

	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func Run() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)
	r.Run(c.Port)
}
