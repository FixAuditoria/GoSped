package Model

import (
	"github.com/jinzhu/gorm"
)

// Reg0100 : Dados do Contabilista
type Reg0100 struct {
	gorm.Model
	Reg      string `gorm:"type:varchar(4)" bson:"reg" json:"reg"`
	Nome     string `gorm:"type:varchar(100)" bson:"nome" json:"nome"`
	Cpf      string `gorm:"type:varchar(11)" bson:"cpf" json:"cpf"`
	Crc      string `gorm:"type:varchar(15)" bson:"crc" json:"crc"`
	Cnpj     string `gorm:"type:varchar(14)" bson:"cnpj" json:"cnpj"`
	Cep      string `gorm:"type:varchar(8)" bson:"cep" json:"cep"`
	End      string `gorm:"type:varchar(60)" bson:"end" json:"end"`
	Num      string `gorm:"type:varchar(10)" bson:"num" json:"num"`
	Compl    string `gorm:"type:varchar(60)" bson:"compl" json:"compl"`
	Bairro   string `gorm:"type:varchar(60)" bson:"bairro" json:"bairro"`
	Fone     string `gorm:"type:varchar(11)" bson:"fone" json:"fone"`
	Fax      string `gorm:"type:varchar(11)" bson:"fax" json:"fax"`
	Email    string `gorm:"type:varchar(100)" bson:"email" json:"email"`
	CodMun   string `gorm:"type:varchar(7)" bson:"codnum" json:"codnum"`
	DtIni    string `gorm:"type:varchar(8)" bson:"dtini" json:"dtini"`
	DtFin    string `gorm:"type:varchar(8)" bson:"dtfin" json:"dtfin"`
	CnpjSped string `gorm:"type:varchar(14)" bson:"cnpjsped" json:"cnpjsped"`
}

// TableName : Funcao responsavel por definir o nome na tabela
func (Reg0100) TableName() string {
	return "reg0100"
}

// CreateDB : funcao para criar o banco de dados do Registro 0100
func (r *Reg0100) CreateDB(db gorm.DB) {
	db.AutoMigrate(r)
}

// DropDB : funcao para apagar banco de dados do Registro 0100
func (r *Reg0100) DropDB(db gorm.DB) {
	db.DropTable(r)
}

//Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0100) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.Nome = l[2]
	r.Cpf = l[3]
	r.Crc = l[4]
	r.Cnpj = l[5]
	r.Cep = l[6]
	r.End = l[7]
	r.Num = l[8]
	r.Compl = l[9]
	r.Bairro = l[10]
	r.Fone = l[11]
	r.Fax = l[12]
	r.Email = l[13]
	r.CodMun = l[14]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
