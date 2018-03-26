package Model

import "time"

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
