package Model

import "time"

type RegC405 struct {
	Reg       string
	DtDoc     time.Time
	Cro       string
	Crz       string
	NumCooFin string
	GtFin     float64
	VlBrt     float64
	DtIni     time.Time
	DtFin     time.Time
	Cnpj      string
}
