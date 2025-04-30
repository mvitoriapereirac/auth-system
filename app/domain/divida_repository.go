package domain

type DividaRepository interface {
	Create(divida Divida, cpf string, empresaID uint) (Divida, error)
	ListByID(id uint) ([]Divida, error)
}
