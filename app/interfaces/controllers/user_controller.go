// package controllers

// import (
//     "github.com/gin-gonic/gin"
//     "auth-system/app/usecases"
//     "net/http"
// )

// type UserController struct {
//     userUsecase usecases.UserUsecase
// }

// func NewUserController(userUsecase usecases.UserUsecase) *UserController {
//     return &UserController{userUsecase}
// }

// func (uc *UserController) GetUser(c *gin.Context) {
//     id := c.Param("id")
//     user, err := uc.userUsecase.GetUserByID(id)
//     if err != nil {
//         c.JSON(http.StatusNotFound, gin.H{"mensagem": "Usuário não encontrado"})
//         return
//     }
//     c.JSON(http.StatusOK, user)
// }


package controllers

import (
    "auth-system/app/usecases"
    "auth-system/app/domain"
    "github.com/gin-gonic/gin"
    "net/http"
)

type UserController struct {
    usecase usecases.UserUsecase
}

// NewUserController cria uma nova instância do UserController
func NewUserController(usecase usecases.UserUsecase) *UserController {
    return &UserController{usecase: usecase}
}

// CreateUser cria um novo usuário
func (ctrl *UserController) CreateUser(c *gin.Context) {
    var user domain.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newUser, err := ctrl.usecase.CreateUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": newUser})
}

func (uc *UserController) Login(c *gin.Context) {
    var loginInput struct {
        Email string `json:"email" binding:"required,email"`
        Senha string `json:"senha" binding:"required"`
    }

    if err := c.ShouldBindJSON(&loginInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }

    token, err := uc.usecase.Login(c.Request.Context(), loginInput.Email, loginInput.Senha)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
