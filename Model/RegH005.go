package Model

import "time"

// RegH005 : Totais do Inventário
type RegH005 struct {
	Reg    string
	DtInv  time.Time
	VlInv  float64
	MotInv string
	DtIni  time.Time
	DtFin  time.Time
	Cnpj   string
}
