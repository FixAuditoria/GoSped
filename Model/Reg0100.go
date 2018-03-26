package Model

import "time"

type Reg0100 struct {
	Reg      string
	Nome     string
	Cpf      string
	Crc      string
	Cnpj     string
	Cep      string
	End      string
	Num      string
	Compl    string
	Bairro   string
	Fone     string
	Fax      string
	Email    string
	CodMun   string
	DtIni    time.Time
	DtFin    time.Time
	CnpjSped string
}
