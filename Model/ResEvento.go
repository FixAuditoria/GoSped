package Model

import "github.com/go-bongo/bongo"

type ResEvento struct {
	bongo.DocumentBase `bson:",inline"`
	COrgao             string `xml:"cOrgao"`
	CNPJ               string `xml:"CNPJ"`
	ChNFe              string `xml:"chNFe"`
	DhEvento           string `xml:"dhEvento"`
	TpEvento           string `xml:"tpEvento"`
	NSeqEvento         string `xml:"nSeqEvento"`
	XEvento            string `xml:"xEvento"`
	DhRecbto           string `xml:"dhRecbto"`
	NProt              string `xml:"nProt"`
}
