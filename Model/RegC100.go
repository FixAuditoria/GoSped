package Model

import "time"

type RegC100 struct {
	Reg        string
	IndOper    string
	IndEmit    string
	CodPart    string
	CodMod     string
	CodSit     string
	Ser        string
	NumDoc     string
	ChvNfe     string
	DtDoc      time.Time
	DtES       time.Time
	VlDoc      float64
	IndPgto    string
	VlDesc     float64
	VlAbatNt   float64
	VlMerc     float64
	IndFrt     string
	VlFrt      float64
	VlSeg      float64
	VlOutDa    float64
	VlBcIcms   float64
	VlIcms     float64
	VlBcIcmsSt float64
	VlIcmsSt   float64
	VlIpi      float64
	VlPis      float64
	VlCofins   float64
	VlPisSt    float64
	VlCofinsSt float64
	DtIni      time.Time
	DtFin      time.Time
	Cnpj       string
}
