package Model

import "time"

type RegH010 struct {
	Reg      string
	CodItem  string
	Unid     string
	Qtd      float64
	VlUnit   float64
	VlItem   float64
	IndProp  string
	CodPart  string
	TxtCompl string
	CodCta   string
	VlItemIr float64
	DtInv    time.Time
	DtIni    time.Time
	DtFin    time.Time
	Cnpj     string
}
