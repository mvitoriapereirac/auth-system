// package routes

// import (
//     "auth-system/app/interfaces/controllers"
//     "github.com/gin-gonic/gin"
// )

// func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
//     router.GET("/ping", func(c *gin.Context) {
//         c.JSON(200, gin.H{"message": "pong"})
//     })
//     router.POST("/register", userController.CreateUser)
// 	router.POST("/login", userController.Login)

// }

package routes

import (
	"auth-system/app/interfaces/controllers"
	"auth-system/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.UserController, dividaController *controllers.DividaController) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	app.Post("/register", userController.CreateUser)
	app.Post("/login", userController.Login)

	// Rotas protegidas por middleware de autenticação
	user := app.Group("/user", middleware.AuthRequired())
	user.Post("/logout", userController.Logout)
	user.Get("/score", dividaController.CalcularScore)
	user.Post("/dividas", dividaController.CriarDivida)
	user.Get("/dividas", dividaController.ListarDividas)

}
