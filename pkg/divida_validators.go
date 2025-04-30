package validators

import (
	"auth-system/app/domain"
	"auth-system/pkg/utils"
	"errors"
	"time"
)

// TODO: refinar validações
func ValidateDividaInput(divida domain.Divida, cpf string) error {

	if !utils.IsValidCPF(cpf) {
		return errors.New("CPF inválido")
	}

	layout := "02-01-2006"
	_, err := time.Parse(layout, divida.Vencimento)

	if err != nil {
		return errors.New("data de vencimento inválida, use o formato DD-MM-AAAA")
	}

	if divida.Valor <= 0 {
		return errors.New("o valor da dívida deve ser maior do que zero")
	}

	return nil
}
