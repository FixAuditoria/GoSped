package Model

import (
	"github.com/jinzhu/gorm"
)

// Reg0200 : Tabela de Identificação do Item (Produtos e Serviços)
type Reg0200 struct {
	gorm.Model
	Reg        string `gorm:"type:varchar(4)" bson:"reg" json:"reg"`
	CodItem    string `gorm:"type:varchar(60)" bson:"coditem" json:"coditem"`
	DescrItem  string `gorm:"type:varchar(200)" bson:"descritem" json:"descritem"`
	CodBarra   string `gorm:"type:varchar(100)" bson:"codbarra" json:"codbarra"`
	CodAntItem string `gorm:"type:varchar(60)" bson:"codantitem" json:"codantitem"`
	UnidInv    string `gorm:"type:varchar(6)" bson:"unidinv" json:"unidinv"`
	TipoItem   string `gorm:"type:varchar(2)" bson:"tipoitem" json:"tipoitem"`
	CodNcm     string `gorm:"type:varchar(8)" bson:"codncm" json:"codncm"`
	ExIpi      string `gorm:"type:varchar(3)" bson:"exipi" json:"exipi"`
	CodGen     string `gorm:"type:varchar(2)" bson:"codgen" json:"codgen"`
	CodLst     string `gorm:"type:varchar(5)" bson:"codlst" json:"codlst"`
	AliqIcms   string `gorm:"type:varchar(10)" bson:"aliqicms" json:"aliqicms"`
	DtIni      string `gorm:"type:varchar(8)" bson:"dtini" json:"dtini"`
	DtFin      string `gorm:"type:varchar(8)" bson:"dtfin" json:"dtfin"`
	CnpjSped   string `gorm:"type:varchar(14)" bson:"cnpjsped" json:"cnpjsped"`
}

// TableName : Funcao responsavel por definir o nome na tabela
func (Reg0200) TableName() string {
	return "reg0200"
}

// CreateDB : funcao para criar o banco de dados do Registro 0100
func (r *Reg0200) CreateDB(db gorm.DB) {
	db.AutoMigrate(r)
}

// DropDB : funcao para apagar banco de dados do Registro 0100
func (r *Reg0200) DropDB(db gorm.DB) {
	db.DropTable(r)
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0200) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodItem = l[2]
	r.DescrItem = l[3]
	r.CodBarra = l[4]
	r.CodAntItem = l[5]
	r.UnidInv = l[6]
	r.TipoItem = l[7]
	r.CodNcm = l[8]
	r.ExIpi = l[9]
	r.CodGen = l[10]
	r.CodLst = l[11]
	r.AliqIcms = l[12]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
