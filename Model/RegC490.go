package Model

import "time"

type RegC490 struct {
	Reg      string
	CstIcms  string
	Cfop     string
	AliqIcms float64
	VlOpr    float64
	VlBcIcms float64
	VlIcms   float64
	CodObs   string
	DtIni    time.Time
	DtFin    time.Time
	Cnpj     string
}
