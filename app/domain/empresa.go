package domain

type Empresa struct {
	ID                  uint   `json:"id"`
	CNPJ                string `json:"cnpj"`
	RazaoSocial         string `json:"razaoSocial"`
	DataHoraUltimoLogin string `json:"dataHoraLogin"`
}
