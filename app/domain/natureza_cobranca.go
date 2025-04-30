package domain

type NaturezaCobranca struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	RazaoCobranca string `json:"razaoCobranca"`
}
