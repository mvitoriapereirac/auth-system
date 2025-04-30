package interfaces

import (
	"auth-system/app/interfaces/controllers"
	"auth-system/app/interfaces/repositories"
	"auth-system/app/interfaces/routes"
	"auth-system/app/services"
	"auth-system/app/usecases"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	app   *fiber.App
	db    *gorm.DB
	redis *redis.Client
}

func NewServer(db *gorm.DB, redisClient *redis.Client) *Server {
	app := fiber.New()

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService()
	userUsecase := usecases.NewUserUsecase(userRepo, authService)
	userController := controllers.NewUserController(userUsecase)

	dividaRepo := repositories.NewDividaRepository(db)
	dividaUsecase := usecases.NewDividaUsecase(dividaRepo, userRepo)
	dividaController := controllers.NewDividaController(dividaUsecase)

	routes.SetupRoutes(app, userController, dividaController)

	return &Server{
		app:   app,
		db:    db,
		redis: redisClient,
	}
}

func (s *Server) Run(addr string) {
	s.app.Listen(addr)
}
