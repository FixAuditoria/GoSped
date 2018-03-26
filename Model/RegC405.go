package Model

import "time"

// RegC405 : Redução Z (código 02, 2D e 60)
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
