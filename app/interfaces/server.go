package interfaces

import (
    "auth-system/app/interfaces/controllers"
    "auth-system/app/interfaces/routes"
    "auth-system/app/interfaces/repositories"
	"auth-system/app/services"
	"auth-system/app/usecases"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "gorm.io/gorm"
)

type Server struct {
    db     *gorm.DB
    redis  *redis.Client
    router *gin.Engine
}

func NewServer(db *gorm.DB, redisClient *redis.Client) *Server {
    router := gin.Default()

    // Inicializa repositórios, usecases e controllers
    userRepo := repositories.NewUserRepository(db)
    // userUsecase := usecases.NewUserUsecase(userRepository)
    // userController := controllers.NewUserController(userUsecase)
	// db := config.InitDB(...)
	// rdb := config.InitRedis(...)

	authService := services.NewAuthService()
	userUsecase := usecases.NewUserUsecase(userRepo, authService)
	userController := controllers.NewUserController(userUsecase)

	// routes.SetupRoutes(router, userController) // passando controller para as rotas


    // Setup de rotas
    routes.SetupRoutes(router, userController)

    return &Server{
        db:     db,
        redis:  redisClient,
        router: router,
    }
}

func (s *Server) Run(addr string) {
    s.router.Run(addr)
}
