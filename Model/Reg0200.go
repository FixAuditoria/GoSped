package Model

import "time"

type Reg0200 struct {
	Reg        string
	CodItem    string
	DescrItem  string
	CodBarra   string
	CodAntItem string
	UnidInv    string
	TipoItem   string
	CodNcm     string
	ExIpi      string
	CodGen     string
	CodLst     string
	AliqIcms   float64
	DtIni      time.Time
	DtFin      time.Time
	Cnpj       string
}
