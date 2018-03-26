package Model

import "time"

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
