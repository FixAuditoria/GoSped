package Model

import "time"

type Reg0220 struct {
	Reg      string
	UnidConv string
	FatConv  float64
	UnidCod  string
	CodItem  string
	DtIni    time.Time
	DtFin    time.Time
	Cnpj     string
	Feito    string
}
