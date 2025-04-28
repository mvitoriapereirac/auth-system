package domain

type Consumidor struct {
	ID                  uint    `json:"id"`
	UserID              uint    `json:"userId"`
	Score               float64 `json:"score"`
	DataHoraUltimoLogin string  `json:"dataHoraLogin"`
}
