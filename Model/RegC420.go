package Model

import "time"

// RegC420 : Registro dos Totalizadores Parciais da Redução Z (código 02, 2D e 60)
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
