package validators

import (
	"auth-system/app/core/constants"
	"auth-system/app/domain"
	"auth-system/pkg/utils"
	"errors"
	"regexp"
	"strings"
	"time"
)

// Validações gerais
func ValidateUserInput(user domain.User) error {

	// CPF
	if !utils.IsValidCPF(user.CPF) {
		return errors.New("CPF inválido")
	}

	// Data de nascimento (ex: 05-07-1999)
	layout := "02-01-2006"
	birthDate, err := time.Parse(layout, user.DataNascimento)
	if err != nil {
		return errors.New("data de nascimento inválida, use o formato DD-MM-AAAA")
	}
	if !isAdult(birthDate) {
		return errors.New("usuário deve ser maior de idade")
	}

	// Email simples
	if !isValidEmail(user.Email) {
		return errors.New("email inválido")
	}

	return nil
}

func isAdult(birthDate time.Time) bool {
	years := time.Since(birthDate).Hours() / 24 / 365.25
	return years >= 18
}

func isValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsEmailAdmin(email string) bool {
	email = strings.TrimSpace(email)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@br\.experian\.com$`)
	return re.MatchString(email)
}

func CheckUserRole(email string, role *string) string {
	isEmailAdmin := IsEmailAdmin(email)

	if isEmailAdmin {
		if role == nil || *role == "" || *role == constants.UserTypeEmpresa {
			return constants.UserTypeEmpresa
		}
		return constants.UserTypeInadequado
	}

	if role == nil || *role == "" || *role == constants.UserTypeConsumidor {
		return constants.UserTypeConsumidor
	}

	return constants.UserTypeInadequado
}
