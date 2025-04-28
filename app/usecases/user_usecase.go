package usecases

import (
    "auth-system/app/domain"
    "auth-system/app/interfaces/repositories"
)

// UserUsecase define os métodos para interagir com o domínio de usuários
type UserUsecase interface {
    CreateUser(user domain.User) (domain.User, error)
    GetUserByID(id uint) (domain.User, error)
}

// userUsecase implementa a interface UserUsecase
type userUsecase struct {
    userRepository domain.UserRepository
}

// NewUserUsecase cria uma nova instância de userUsecase
func NewUserUsecase(repo domain.UserRepository) UserUsecase {
    return &userUsecase{userRepository: repo}
}

// CreateUser cria um novo usuário no banco de dados
func (uc *userUsecase) CreateUser(user domain.User) (domain.User, error) {
    return uc.userRepository.Create(user)
}

// GetUserByID retorna um usuário pelo ID
func (uc *userUsecase) GetUserByID(id uint) (domain.User, error) {
    return uc.userRepository.GetByID(id)
}
