package auth

import (
	"API_Gateway/configs"
	"API_Gateway/internal/auth/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, c *configs.Config) *ServiceClient {
	svc := &ServiceClient{Client: InitServiceClient(c)}

	routes := r.Group("/")
	routes.POST("sign-up", svc.SignUp)
	routes.POST("sign-in", svc.SignIn)

	return svc
}

func (svc *ServiceClient) SignUp(ctx *gin.Context) {
	routes.SignUp(ctx, svc.Client)
}

func (svc *ServiceClient) SignIn(ctx *gin.Context) {
	routes.SignIn(ctx, svc.Client)
}
