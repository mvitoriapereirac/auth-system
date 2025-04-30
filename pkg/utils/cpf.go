// pkg/utils/cpf.go
package utils

import (
	"regexp"
)

// IsValidCPF checks if a CPF is valid (only basic format and length check)
func IsValidCPF(cpf string) bool {
	// Remove all non-digit characters
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return false
	}

	// TODO: Implement full CPF validation algorithm (check digits)
	return true
}
