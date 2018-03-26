package Model

import "time"

// RegC465 : Complemento do Cupom Fiscal Eletrônico Emitido por ECF - CF-e-ECF (código 60)
type RegC465 struct {
	Reg    string
	ChvCfe string
	NumCcF string
	DtIni  time.Time
	DtFin  time.Time
	Cnpj   string
}
