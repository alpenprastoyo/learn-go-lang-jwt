package controller

import "github.com/gin-gonic/gin"

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	//di sini masukkan service yang kalian butuh
	//this is where you put your service
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ct)