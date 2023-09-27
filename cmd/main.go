package main

import (
	"API_Gateway/configs"
	"API_Gateway/internal/auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.InitConfig()

	r := gin.Default()
	authSvc := *auth.InitRoutes(r, &cfg)
	fmt.Println(authSvc)

	r.Run(cfg.Port)
}
