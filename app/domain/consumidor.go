package domain

type Consumidor struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	UserID uint    `gorm:"not null" json:"userId"`
	Score  float64 `json:"score"`
	User   User    `gorm:"foreignKey:UserID"`
}
