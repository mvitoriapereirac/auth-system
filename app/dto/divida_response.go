package dto

type DividaResponse struct {
	ID         uint    `json:"id"`
	Valor      float64 `json:"valor"`
	Vencimento string  `json:"vencimento"`
}
