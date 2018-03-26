package Model

import "time"

// RegC425 : Resumo de itens do movimento diário (código 02 e 2D)
type RegC425 struct {
	Reg      string
	CodItem  string
	Qtd      float64
	Unid     string
	VlItem   float64
	VlPis    float64
	VlCofins float64
	DtIni    time.Time
	DtFin    time.Time
	Cnpj     string
}
