package domain

type Empresa struct {
    ID                  uint   `gorm:"primaryKey" json:"id"`
    UserID              uint   `gorm:"not null" json:"userId"`
    CNPJ                string `gorm:"unique" json:"cnpj"`
    RazaoSocial         string `json:"razaoSocial"`
    User                User   `gorm:"foreignKey:UserID"`
}