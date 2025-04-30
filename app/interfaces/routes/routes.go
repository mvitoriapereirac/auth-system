package routes

import (
    "auth-system/app/interfaces/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    router.POST("/register", userController.CreateUser)
	router.POST("/login", userController.Login)

}
