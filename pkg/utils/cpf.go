package utils

import (
	"regexp"
)

func IsValidCPF(cpf string) bool {
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return false
	}

	// TODO: Implementar validação completa de CPF
	return true
}
