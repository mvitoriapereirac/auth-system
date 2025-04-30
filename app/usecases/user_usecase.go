package usecases

import (
	"auth-system/app/core/constants"
	"auth-system/app/domain"
	"auth-system/app/dto"
	"auth-system/app/services"
	"auth-system/pkg"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(user domain.User) (dto.UserResponse, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type userUsecase struct {
	userRepository domain.UserRepository
	authService    services.AuthService
}

func NewUserUsecase(ur domain.UserRepository, as services.AuthService) UserUsecase {
	return &userUsecase{
		userRepository: ur,
		authService:    as,
	}
}

func (uc *userUsecase) CreateUser(user domain.User) (dto.UserResponse, error) {
	if err := validators.ValidateUserInput(user); err != nil {
		return dto.UserResponse{}, err
	}

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

func (u *userUsecase) Login(ctx context.Context, cpf, password string) (string, error) {
	user, err := u.userRepository.GetByCPF(ctx, cpf)
	if err != nil {
		return "", errors.New("CPF não encontrado")
	}

	if !u.authService.VerifyPassword(user.Senha, password) {
		return "", errors.New("Senha inválida")
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
