package Model

import "time"

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
