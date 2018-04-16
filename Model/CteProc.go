package Model

import "github.com/go-bongo/bongo"

type CteProc struct {
	bongo.DocumentBase `bson:",inline"`
	Cte                struct {
		InfCte struct {
			Versao string `xml:"versao,attr"`
			Id     string `xml:"Id,attr"`
			Ide    struct {
				CUf       string `xml:"cUF"`
				CCT       string `xml:"cCT"`
				CFOP      string `xml:"CFOP"`
				NatOp     string `xml:"natOp"`
				Mod       string `xml:"mod"`
				Serie     string `xml:"serie"`
				NCT       string `xml:"nCT"`
				DhEmi     string `xml:"dhEmi"`
				TpImp     string `xml:"tpImp"`
				TpEmis    string `xml:"tpEmis"`
				CDV       string `xml:"cDV"`
				TpAmb     string `xml:"tpAmb"`
				TpCTe     string `xml:"tpCTe"`
				ProcEmi   string `xml:"procEmi"`
				VerProc   string `xml:"verProc"`
				CMunEnv   string `xml:"cMunEnv"`
				XMunEnv   string `xml:"xMunEnv"`
				UFEnv     string `xml:"UFEnv"`
				Modal     string `xml:"modal"`
				TpServ    string `xml:"tpServ"`
				CMunIni   string `xml:"cMunIni"`
				XMunIni   string `xml:"xMunIni"`
				UFIni     string `xml:"UFIni"`
				CMunFim   string `xml:"cMunFim"`
				XMunFim   string `xml:"xMunFim"`
				UFFim     string `xml:"UFFim"`
				Retira    string `xml:"retira"`
				IndIEToma string `xml:"indIEToma"`
				Toma3     struct {
					Toma string `xml:"toma"`
				} `xml:"toma3"`
			} `xml:"ide"`
			Compl struct {
				XCaracAd  string `xml:"xCaracAd"`
				XCaracSer string `xml:"xCaracSer"`
				XEmi      string `xml:"xEmi"`
				Entrega   struct {
					NoPeriodo struct {
						TpPer string `xml:"tpPer"`
						DIni  string `xml:"dIni"`
						DFim  string `xml:"dFim"`
					} `xml:"noPeriodo"`
					ComHora struct {
						TpHor string `xml:"tpHor"`
						HProg string `xml:"hProg"`
					} `xml:"comHora"`
				} `xml:"Entrega"`
				Xobs    string `xml:"xObs"`
				ObsCont struct {
					XTexto string `xml:"xTexto"`
				}
			} `xml:"compl"`
			Emit struct {
				CNPJ      string `xml:"CNPJ"`
				IE        string `xml:"IE"`
				XNome     string `xml:"xNome"`
				XFant     string `xml:"xFant"`
				EnderEmit struct {
					XLgr    string `xml:"xLgr"`
					Nro     string `xml:"nro"`
					XBairro string `xml:"xBairro"`
					CMun    string `xml:"cMun"`
					XMun    string `xml:"xMun"`
					CEP     string `xml:"CEP"`
					UF      string `xml:"UF"`
					Fone    string `xml:"fone"`
				} `xml:"enderEmit"`
			} `xml:"emit"`
			Rem struct {
				CNPJ      string `xml:"CNPJ"`
				IE        string `xml:"IE"`
				XNome     string `xml:"xNome"`
				Fone      string `xml:"fone"`
				EnderReme struct {
					XLgr    string `xml:"xLgr"`
					Nro     string `xml:"nro"`
					XBairro string `xml:"xBairro"`
					CMun    string `xml:"cMun"`
					XMun    string `xml:"xMun"`
					CEP     string `xml:"CEP"`
					UF      string `xml:"UF"`
				} `xml:"enderReme"`
			} `xml:"rem"`
			Dest struct {
				CNPJ      string `xml:"CNPJ"`
				IE        string `xml:"IE"`
				XNome     string `xml:"xNome"`
				Fone      string `xml:"fone"`
				EnderDest struct {
					XLgr    string `xml:"xLgr"`
					Nro     string `xml:"nro"`
					XBairro string `xml:"xBairro"`
					CMun    string `xml:"cMun"`
					XMun    string `xml:"xMun"`
					CEP     string `xml:"CEP"`
					UF      string `xml:"UF"`
				} `xml:"enderDest"`
			} `xml:"dest"`
			VPrest struct {
				VTPrest string `xml:"vTPrest"`
				VRec    string `xml:"vRec"`
				Comp    []struct {
					XNome string `xml:"vTPrest"`
					VComp string `xml:"vRec"`
				} `xml:"Comp"`
			} `xml:"vPrest"`
		} `xml:"infCte"`
	} `xml:"CTe"`
}
