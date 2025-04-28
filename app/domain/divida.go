package domain

type Divida struct {
	ID                 uint    `json:"id"`
	Valor              float64 `json:"valor"`
	Vencimento         string  `json:"vencimento"`
	NaturezaCobrancaID uint    `json:"naturezaCobrancaId"`
}
