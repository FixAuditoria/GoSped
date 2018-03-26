package Model

import "time"

// Reg0220 : Fatores de Conversão de Unidades
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
