package controllers

import (
    "github.com/gin-gonic/gin"
    "auth_project_base/app/usecases"
    "net/http"
)

type UserController struct {
    userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) *UserController {
    return &UserController{userUsecase}
}

func (uc *UserController) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := uc.userUsecase.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"mensagem": "Usuário não encontrado"})
        return
    }
    c.JSON(http.StatusOK, user)
}
