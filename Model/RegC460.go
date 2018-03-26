package Model

import "time"

// RegC460 : Documento Fiscal Emitido por ECF (código 02, 2D e 60)
type RegC460 struct {
	Reg      string
	CodMod   string
	CodSit   string
	NumDoc   string
	DtDoc    time.Time
	VlDoc    float64
	VlPis    float64
	VlCofins float64
	CpfCnpj  string
	NomAdq   string
	DtIni    time.Time
	DtFin    time.Time
	Cnpj     string
}
