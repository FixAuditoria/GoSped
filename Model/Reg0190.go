package Model

import (
	"github.com/jinzhu/gorm"
)

// Reg0190 : Identificação das unidades de medida
type Reg0190 struct {
	gorm.Model
	Reg      string `gorm:"type:varchar(4)" bson:"reg" json:"reg"`
	Unid     string `gorm:"type:varchar(6)" bson:"unid" json:"unid"`
	Descr    string `gorm:"type:varchar(150)" bson:"descr" json:"descr"`
	DtIni    string `gorm:"type:varchar(8)" bson:"dtini" json:"dtini"`
	DtFin    string `gorm:"type:varchar(8)" bson:"dtfin" json:"dtfin"`
	CnpjSped string `gorm:"type:varchar(14)" bson:"cnpjsped" json:"cnpjsped"`
}

// TableName : Funcao responsavel por definir o nome na tabela
func (Reg0190) TableName() string {
	return "reg0190"
}

// CreateDB : funcao para criar o banco de dados do Registro 0100
func (r *Reg0190) CreateDB(db gorm.DB) {
	db.AutoMigrate(r)
}

// DropDB : funcao para apagar banco de dados do Registro 0100
func (r *Reg0190) DropDB(db gorm.DB) {
	db.DropTable(r)
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0190) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.Unid = l[2]
	r.Descr = l[3]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
