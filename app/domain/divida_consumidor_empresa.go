package domain

type DividaConsumidorEmpresa struct {
	ID uint `json:"id"`
	ConsumidorID uint `json:"consumidorId"`
    DividaID     uint `json:"dividaId"`
    EmpresaID    uint `json:"empresaId"`
    FoiVisualizada bool `json:"foiVisualizada"`
}