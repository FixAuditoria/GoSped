package Model

type CteProc struct {
	Cte struct {
		InfCte struct {
			Ide struct {
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
					// #### PAREI AQUI ###########
				} `xml:"Entrega"`
			} `xml:"compl"`
		} `xml:"infeCte"`
	} `xml:"CTe"`
}
