package Model

import "time"

type RegC420 struct {
	Reg        string
	CodTotPar  string
	VlrAcumTot float64
	NrTot      string
	DescrNrTot string
	DtIni      time.Time
	DtFin      time.Time
	Cnpj       string
}
