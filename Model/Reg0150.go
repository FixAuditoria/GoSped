package Model

import "time"

// Reg0150 : Tabela de Cadastro do Participante
type Reg0150 struct {
	Reg      string
	CodPart  string
	Nome     string
	CodPais  string
	Cnpj     string
	Cpf      string
	Ie       string
	CodMun   string
	Suframa  string
	Endereco string
	Num      string
	Compl    string
	Bairro   string
	DtIni    time.Time
	DtFin    time.Time
	CnpjSped string
}
