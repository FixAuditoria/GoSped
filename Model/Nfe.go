package Model

import "github.com/go-bongo/bongo"

type NFe struct {
	bongo.DocumentBase `bson:",inline"`
	InfNFe             struct {
		Versao string `xml:"versao,attr"`
		Id     string `xml:"Id,attr"`
		Ide    struct {
			Cuf      string `xml:"cUF"`
			Cnf      string `xml:"cNF"`
			NatOp    string `xml:"natOp"`
			IndPag   string `xml:"indPag"`
			Mod      string `xml:"mod"`
			Serie    string `xml:"serie"`
			Nnf      string `xml:"nNF"`
			DhEmit   string `xml:"dhEmi"`
			DhSaiEnt string `xml:"dhSaiEnt"`
			TpNf     string `xml:"tpNF"`
			IdDest   string `xml:"idDest"`
			CmunFg   string `xml:"cMunFG"`
			TpImp    string `xml:"tpImp"`
			TpEmis   string `xml:"tpEmis"`
			Cdv      string `xml:"cDV"`
			TpAmb    string `xml:"tpAmb"`
			FinNfe   string `xml:"finNFe"`
			IndFinal string `xml:"indFinal"`
			IndPres  string `xml:"indPres"`
			ProcEmi  string `xml:"procEmi"`
			VerProc  string `xml:"verProc"`
		} `xml:"ide"`
		Emit struct {
			Cnpj      string `xml:"CNPJ"`
			Xnome     string `xml:"xNome"`
			Xfant     string `xml:"xFant"`
			EnderEmit struct {
				XLgr    string `xml:"xLgr"`
				Nro     string `xml:"nro"`
				Xbairro string `xml:"xBairro"`
				Cmun    string `xml:"cMun"`
				Xmun    string `xml:"xMun"`
				Uf      string `xml:"UF"`
				Cep     string `xml:"CEP"`
				Cpais   string `xml:"cPais"`
				Xpais   string `xml:"xPais"`
				Fone    string `xml:"fone"`
			} `xml:"enderEmit"`
			Ie  string `xml:"IE"`
			Crt string `xml:"CRT"`
		} `xml:"emit"`
		Dest struct {
			Cnpj      string `xml:"CNPJ"`
			Xnome     string `xml:"xNome"`
			Xfant     string `xml:"xFant"`
			EnderDest struct {
				XLgr    string `xml:"xLgr"`
				Nro     string `xml:"nro"`
				Xbairro string `xml:"xBairro"`
				Cmun    string `xml:"cMun"`
				Xmun    string `xml:"xMun"`
				Uf      string `xml:"UF"`
				Cep     string `xml:"CEP"`
				Cpais   string `xml:"cPais"`
				Xpais   string `xml:"xPais"`
				Fone    string `xml:"fone"`
			} `xml:"enderDest"`
			Ie  string `xml:"IE"`
			Crt string `xml:"CRT"`
		} `xml:"dest"`
		Det []struct {
			NItem string `xml:"nItem,attr"`
			Prod  struct {
				CProd    string `xml:"cProd"`
				CEAN     string `xml:"cEAN"`
				XProd    string `xml:"xProd"`
				NCM      string `xml:"NCM"`
				CFOP     string `xml:"CFOP"`
				UCom     string `xml:"uCom"`
				QCom     string `xml:"qCom"`
				VUnCom   string `xml:"vUnCom"`
				VProd    string `xml:"vProd"`
				CEANTrib string `xml:"cEANTrib"`
				UTrib    string `xml:"uTrib"`
				QTrib    string `xml:"qTrib"`
				VUnTrib  string `xml:"vUnTrib"`
				IndTot   string `xml:"indTot"`
			} `xml:"prod"`
			Imposto struct {
				VtotTrib string `xml:"vTotTrib"`
				Icms     struct {
					Icms00 struct {
						Orig string `xml:"orig"`
						Cst  string `xml:"CST"`
					} `xml:"ICMS00"`
					Icms60 struct {
						Orig string `xml:"orig"`
						Cst  string `xml:"CST"`
					} `xml:"ICMS60"`
				} `xml:"ICMS"`
			} `xml:"imposto"`
		} `xml:"det"`
	} `xml:"infNFe"`
}
