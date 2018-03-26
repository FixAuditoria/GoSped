package Model

import "time"

// RegC400 : Equipamento ECF (código 02, 2D e 60),
type RegC400 struct {
	Reg    string
	CodMod string
	EcfMod string
	EcfFab string
	EcfCx  string
	DtIni  time.Time
	DtFin  time.Time
	Cnpj   string
}
