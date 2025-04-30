package domain

type DividaConsumidorEmpresa struct {
    ID            uint `gorm:"primaryKey" json:"id"`
    ConsumidorID  uint `gorm:"not null" json:"consumidorId"`
    DividaID      uint `gorm:"not null" json:"dividaId"`
    EmpresaID     uint `gorm:"not null" json:"empresaId"`
    FoiVisualizada bool `json:"foiVisualizada"`
}