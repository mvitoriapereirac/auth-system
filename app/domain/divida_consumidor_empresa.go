package domain

// A estrutura DividaConsumidorEmpresa representa a relação entre dívida, consumidor e empresa.
// Embora, para o escopo atual do projeto (um teste técnico com regras fixas e fluxo fechado),
// fosse possível armazenar diretamente os metadados de consumidor e empresa na própria tabela de dívida,
// optei por modelar a relação como uma tabela intermediária para demonstrar uma visão mais escalável.
//
// Essa escolha antecipa possíveis evoluções no sistema, como a transferência de dívidas entre empresas,
// múltiplos consumidores vinculados a uma única dívida (como fiadores), e a necessidade de manter histórico
// de interações ou negociações. Mesmo que não seja utilizada em sua totalidade neste momento,
// essa estrutura simboliza um pensamento voltado à robustez e flexibilidade do domínio.
type DividaConsumidorEmpresa struct {
	ID             uint `gorm:"primaryKey" json:"id"`
	ConsumidorID   uint `gorm:"not null" json:"consumidorId"`
	DividaID       uint `gorm:"not null" json:"dividaId"`
	EmpresaID      uint `gorm:"not null" json:"empresaId"`
	FoiVisualizada bool `json:"foiVisualizada"`
}
