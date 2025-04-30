package domain

type User struct {
    ID             uint   `gorm:"primaryKey" json:"id"`
    Tipo           *string `gorm:"type:char(1);not null" json:"tipo"`
    CPF            string `gorm:"unique;not null" json:"cpf"`
    DataNascimento string `gorm:"type:date;not null" json:"dataNascimento"`
    Email          string `gorm:"unique;not null" json:"email"`
    Senha          string `gorm:"not null" json:"senha"`
}
