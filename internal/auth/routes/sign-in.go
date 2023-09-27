package routes

import (
	"API_Gateway/internal/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignInRequestBody struct {
	PhoneNumber int64  `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func SignIn(ctx *gin.Context, client pb.AuthServiceClient) {
	b := &SignInRequestBody{}

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.SignIn(ctx, &pb.SignInRequest{
		PhoneNumber: b.PhoneNumber,
		Password:    b.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
