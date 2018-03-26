package Model

import "time"

// Reg0190 : Identificação das unidades de medida
type Reg0190 struct {
	Reg   string
	Unid  string
	Descr string
	DtIni time.Time
	DtFin time.Time
	Cnpj  string
}
