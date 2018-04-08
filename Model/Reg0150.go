package Model

import (
	"github.com/jinzhu/gorm"
)

// Reg0150 : Tabela de Cadastro do Participante
type Reg0150 struct {
	gorm.Model
	Reg      string `gorm:"type:varchar(4)" bson:"reg" json:"reg"`
	CodPart  string `gorm:"type:varchar(60)" bson:"codpart" json:"codpart"`
	Nome     string `gorm:"type:varchar(100)" bson:"nome" json:"nome"`
	CodPais  string `gorm:"type:varchar(5)" bson:"codpais" json:"codpais"`
	Cnpj     string `gorm:"type:varchar(14)" bson:"cnpj" json:"cnpj"`
	Cpf      string `gorm:"type:varchar(11)" bson:"cpf" json:"cpf"`
	Ie       string `gorm:"type:varchar(14)" bson:"ie" json:"ie"`
	CodMun   string `gorm:"type:varchar(7)" bson:"codmun" json:"codmun"`
	Suframa  string `gorm:"type:varchar(9)" bson:"suframa" json:"suframa"`
	Endereco string `gorm:"type:varchar(60)" bson:"endereco" json:"endereco"`
	Num      string `gorm:"type:varchar(10)" bson:"num" json:"num"`
	Compl    string `gorm:"type:varchar(60)" bson:"compl" json:"compl"`
	Bairro   string `gorm:"type:varchar(60)" bson:"bairro" json:"bairro"`
	DtIni    string `gorm:"type:varchar(8)" bson:"dtini" json:"dtini"`
	DtFin    string `gorm:"type:varchar(8)" bson:"dtfin" json:"dtfin"`
	CnpjSped string `gorm:"type:varchar(14)" bson:"cnpjsped" json:"cnpjsped"`
}

// TableName : Funcao responsavel por definir o nome na tabela
func (Reg0150) TableName() string {
	return "reg0150"
}

// CreateDB : funcao para criar o banco de dados do Registro 0100
func (r *Reg0150) CreateDB(db gorm.DB) {
	db.AutoMigrate(r)
}

// DropDB : funcao para apagar banco de dados do Registro 0100
func (r *Reg0150) DropDB(db gorm.DB) {
	db.DropTable(r)
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0150) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodPart = l[2]
	r.Nome = l[3]
	r.CodPais = l[4]
	r.Cnpj = l[5]
	r.Cpf = l[6]
	r.Ie = l[7]
	r.CodMun = l[8]
	r.Suframa = l[9]
	r.Endereco = l[10]
	r.Num = l[11]
	r.Compl = l[12]
	r.Bairro = l[13]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
