package controller


//AuthController interface is a contract what this controller can do
type AuthController inteface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}