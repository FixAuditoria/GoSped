package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// ProdutosNfe : Produtos da base auditoria estraido via scrapy pelo sistema NfeScrapy
type ProdutosNfe struct {
	ID               bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Chave            string        `bson:"chave" json:"chave"`
	Sequencial       string        `bson:"sequencial" json:"sequencial"`
	Descricao        string        `bson:"descricao" json:"descricao"`
	Qtd              string        `bson:"qtd" json:"qtd"`
	Un               string        `bson:"un" json:"un"`
	VlTot            string        `bson:"vlTot" json:"vlTot"`
	Codigo           string        `bson:"codigo" json:"codigo"`
	Ncm              string        `bson:"ncm" json:"ncm"`
	Cest             string        `bson:"cest" json:"cest"`
	ExTipi           string        `bson:"exTipi" json:"exTipi"`
	Cfop             string        `bson:"cfop" json:"cfop"`
	OutrasDespesas   string        `bson:"outrasDespesas" json:"outrasDespesas"`
	VlDesc           string        `bson:"vlDesc" json:"vlDesc"`
	VlTotFrete       string        `bson:"vlTotFrete" json:"vlTotFrete"`
	VlSeguro         string        `bson:"vlSeguro" json:"vlSeguro"`
	IndComposicao    string        `bson:"indComposicao" json:"indComposicao"`
	EanComercial     string        `bson:"eanComercial" json:"eanComercial"`
	UnComercial      string        `bson:"unComercial" json:"unComercial"`
	QtdComercial     string        `bson:"qtdComercial" json:"qtdComercial"`
	EanTributavel    string        `bson:"eanTributavel" json:"eanTributavel"`
	UnTributavel     string        `bson:"unTributavel" json:"unTributavel"`
	QtdTributavel    string        `bson:"qtdTributavel" json:"qtdTributavel"`
	VlUnitComercial  string        `bson:"vlUnitComercial" json:"vlUnitComercial"`
	VlUnitTributavel string        `bson:"vlUnitTributavel" json:"vlUnitTributavel"`
	NumPedido        string        `bson:"numPedido" json:"numPedido"`
	ItemPedido       string        `bson:"itemPedido" json:"itemPedido"`
	VlAproxTributo   string        `bson:"vlAproxTributo" json:"vlAproxTributo"`
	NumFci           string        `bson:"numFci" json:"numFci"`
}
