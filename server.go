package main

import (
	"github.com/alpenprastoyo/learn-go-lang-jwt/config"
	"github.com/alpenprastoyo/learn-go-lang-jwt/controller"
	"github.com/alpenprastoyo/learn-go-lang-jwt/repository"
	"github.com/alpenprastoyo/learn-go-lang-jwt/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	// , middleware.AutorizeJWT(jwtService)

	r.Run()

}
