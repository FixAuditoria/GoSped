package Model

type ProcEventoNFe struct {
	Versao string `xml:"versao,attr"`
	Evento struct {
		InfEvento struct {
			Id         string `xml:"Id,attr"`
			COrgao     string `xml:"cOrgao"`
			TpAmb      string `xml:"tpAmb"`
			CNPJ       string `xml:"CNPJ"`
			ChNFe      string `xml:"chNFe"`
			HhEvento   string `xml:"dhEvento"`
			TpEvento   string `xml:"tpEvento"`
			NSeqEvento string `xml:"nSeqEvento"`
			VerEvento  string `xml:"verEvento"`
			DetEvento  struct {
				Versao     string `xml:"versao,attr"`
				DescEvento string `xml:"descEvento"`
				NProt      string `xml:"nProt"`
				XJust      string `xml:"xJust"`
			} `xml:"detEvento"`
		} `xml:"infEvento"`
	} `xml:"evento"`
	RetEvento struct {
		InfEvento struct {
			TpAmb       string `xml:"tpAmb"`
			VerAplic    string `xml:"verAplic"`
			COrgao      string `xml:"cOrgao"`
			CStat       string `xml:"cStat"`
			XMotivo     string `xml:"xMotivo"`
			ChNFe       string `xml:"chNFe"`
			TpEvento    string `xml:"tpEvento"`
			XEvento     string `xml:"xEvento"`
			NSeqEvento  string `xml:"nSeqEvento"`
			DhRegEvento string `xml:"dhRegEvento"`
			NProt       string `xml:"nProt"`
		} `xml:"infEvento"`
	} `xml:"retEvento"`
}
