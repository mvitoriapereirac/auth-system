package dto

type UserResponse struct {
	ID             uint   `json:"id"`
	Tipo           string `json:"tipo"`
	CPF            string `json:"cpf"`
	DataNascimento string `json:"dataNascimento"`
	Email          string `json:"email"`
}
