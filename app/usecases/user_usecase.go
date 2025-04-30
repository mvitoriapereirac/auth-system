package usecases

// import (
//     "auth-system/app/domain"
//     "auth-system/app/interfaces/repositories"
// )

// // UserUsecase define os métodos para interagir com o domínio de usuários
// type UserUsecase interface {
//     CreateUser(user domain.User) (domain.User, error)
//     GetUserByID(id uint) (domain.User, error)
// }

// // userUsecase implementa a interface UserUsecase
// type userUsecase struct {
//     userRepository domain.UserRepository
// }

// // NewUserUsecase cria uma nova instância de userUsecase
// func NewUserUsecase(repo domain.UserRepository) UserUsecase {
//     return &userUsecase{userRepository: repo}
// }

// // CreateUser cria um novo usuário no banco de dados
// func (uc *userUsecase) CreateUser(user domain.User) (domain.User, error) {
//     return uc.userRepository.Create(user)
// }

// // GetUserByID retorna um usuário pelo ID
// func (uc *userUsecase) GetUserByID(id uint) (domain.User, error) {
//     return uc.userRepository.GetByID(id)
// }

// package usecases

import (
    "context"
    "errors"
	"auth-system/app/dto"
    "auth-system/app/domain"
    "auth-system/app/services"
	"golang.org/x/crypto/bcrypt"
	"auth-system/pkg"
	"auth-system/app/core/constants"

)

// UserUsecase define os métodos que vão orquestrar a lógica de negócios dos usuários
type UserUsecase interface {
    CreateUser(user domain.User) (dto.UserResponse, error)
    Login(ctx context.Context, email, password string) (string, error)
}

// userUsecase implementa a interface UserUsecase
type userUsecase struct {
    userRepository domain.UserRepository
    authService    services.AuthService
}

// NewUserUsecase cria uma nova instância de userUsecase
func NewUserUsecase(ur domain.UserRepository, as services.AuthService) UserUsecase {
    return &userUsecase{
        userRepository: ur,
        authService:    as,
    }
}

// CreateUser cria um novo usuário
func (uc *userUsecase) CreateUser(user domain.User) (dto.UserResponse, error) {
	if err := validators.ValidateUserInput(user); err != nil {
		return dto.UserResponse{}, err
	}

	//Checa validez de tipo
	tipo := validators.CheckUserRole(user.Email, user.Tipo)
	if tipo == constants.UserTypeInadequado {
		return dto.UserResponse{}, errors.New("O tipo de usuário está errado. Por favor, adeque o domínio do seu email ou envie outro tipo.")
	}
	user.Tipo = &tipo
	
	if uc.userRepository.ExistsByCPF(user.CPF) {
		return dto.UserResponse{}, errors.New("CPF já cadastrado")
	}
	
	if uc.userRepository.ExistsByEmail(user.Email) {
		return dto.UserResponse{}, errors.New("email já cadastrado")
	}

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Senha), bcrypt.DefaultCost)
    if err != nil {
        return dto.UserResponse{}, err
    }
    user.Senha = string(hashedPassword)
    createdUser, err := uc.userRepository.CreateWithRole(user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	userResponse := dto.UserResponse{
		ID:             createdUser.ID,
		Tipo:           *createdUser.Tipo,
		CPF:            createdUser.CPF,
		DataNascimento: createdUser.DataNascimento,
		Email:          createdUser.Email,
	}
	return userResponse, nil

}

// Login realiza o login de um usuário
func (u *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
    user, err := u.userRepository.GetByEmail(ctx, email)
    if err != nil {
        return "", errors.New("usuário ou senha inválidos")
    }

    if !u.authService.VerifyPassword(user.Senha, password) {
        return "", errors.New("usuário ou senha inválidos")
    }

    token, err := u.authService.GenerateToken(user.ID)
    if err != nil {
        return "", err
    }

    err = u.authService.SaveToken(ctx, user.ID, token)
    if err != nil {
        return "", err
    }

    return token, nil
}
