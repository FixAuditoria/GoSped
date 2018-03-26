package Model

import "time"

type RegC170 struct {
	Reg           string
	NumItem       string
	CodItem       string
	DescrCompl    string
	Qtd           float64
	Unid          string
	VlItem        float64
	VlDesc        float64
	IndMov        string
	CstIcms       string
	Cfop          string
	CodNat        string
	VlBcIcms      float64
	AliqIcms      float64
	VlIcms        float64
	VlBcIcmsSt    float64
	AliqSt        float64
	VlIcmsSt      float64
	IndApur       string
	CstIpi        string
	CodEnq        string
	VlBcIpi       float64
	AliqIpi       float64
	VlIpi         float64
	CstPis        string
	VlBcPis       float64
	AliqPis01     float64
	QuantBcPis    float64
	AliqPis02     float64
	VlPis         float64
	CstCofins     string
	VlBcCofins    float64
	AliqCofins01  float64
	QuantBcCofins float64
	AliqCofins02  float64
	VlCofins      float64
	CodCta        string
	EntradaSaida  string // Se for entrada 0, se for saida 1
	NumDoc        string
	DtIni         time.Time
	DtFin         time.Time
	Cnpj          string
}
