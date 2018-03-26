package Model

import "time"

// RegC190 : Registro Analítico do Documento (código 01, 1B, 04, 55 e 65)
type RegC190 struct {
	Reg        string
	CstIcms    string
	AliqIcms   string
	VlOpr      string
	VlBcIcms   string
	VlIcms     string
	VlBcIcmsSt string
	VlIcmsSt   string
	VlRedBc    string
	VlIpi      string
	CodObs     string
	DtIni      time.Time
	DtFin      time.Time
	CnpjSped   string
}
