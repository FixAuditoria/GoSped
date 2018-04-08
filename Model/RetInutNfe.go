package Model

import "github.com/go-bongo/bongo"

type RetInutNfe struct {
	bongo.DocumentBase `bson:",inline"`
	Xmlns              string `xml:"xmlns,attr"`
	Versao             string `xml:"versao,attr"`
	InfInut            struct {
		TpAmb    string `xml:"tpAmb"`
		VerAplic string `xml:"verAplic"`
		CStat    string `xml:"cStat"`
		XMotivo  string `xml:"xMotivo"`
		CUF      string `xml:"cUF"`
		Ano      string `xml:"ano"`
		CNPJ     string `xml:"CNPJ"`
		Mod      string `xml:"mod"`
		Serie    string `xml:"serie"`
		NNFIni   string `xml:"nNFIni"`
		NNFFin   string `xml:"nNFFin"`
		DhRecbto string `xml:"dhRecbto"`
		NProt    string `xml:"nProt"`
	} `xml:"infInut"`
}
