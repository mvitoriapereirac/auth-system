package domain

type Divida struct {
    ID                 uint    `gorm:"primaryKey" json:"id"`
    Valor              float64 `gorm:"not null" json:"valor"`
    Vencimento         string  `gorm:"type:date;not null" json:"vencimento"`
    NaturezaCobrancaID uint    `json:"naturezaCobrancaId"`
    NaturezaCobranca   NaturezaCobranca `gorm:"foreignKey:NaturezaCobrancaID"`
}
