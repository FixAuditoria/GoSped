package Model

import "time"

// RegC470 : Itens do Documento Fiscal Emitido por ECF (código 02 e 2D)
type RegC470 struct {
	Reg      string
	CodItem  string
	Qtd      float64
	QtdCanc  float64
	Unid     string
	VlItem   float64
	CstIcms  string
	Cfop     string
	AliqIcms float64
	VlPis    float64
	VlCofins float64
	DtIni    time.Time
	DtFin    time.Time
	Cnpj     string
}
