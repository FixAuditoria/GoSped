package Model

import "github.com/go-bongo/bongo"

type ProcEventoCTe struct {
	bongo.DocumentBase `bson:",inline"`
	EventoCTe          struct {
		InfEvento struct {
			COrgao     string `xml:"cOrgao"`
			TpAmb      string `xml:"tpAmb"`
			CNPJ       string `xml:"CNPJ"`
			ChCTe      string `xml:"chCTe"`
			DhEvento   string `xml:"dhEvento"`
			TpEvento   string `xml:"tpEvento"`
			NSeqEvento string `xml:"nSeqEvento"`
			DetEvento  struct {
				EvCTeAutorizadoMDFe struct {
					DescEvento string `xml:"descEvento"`
					MDFe       struct {
						ChMDFe   string `xml:"chMDFe"`
						Modal    string `xml:"modal"`
						DhEmi    string `xml:"dhEmi"`
						NProt    string `xml:"nProt"`
						DhRecbto string `xml:"dhRecbto"`
					} `xml:"MDFe"`
					Emit struct {
						CNPJ  string `xml:"CNPJ"`
						IE    string `xml:"IE"`
						XNome string `xml:"xNome"`
					} `xml:"emit"`
				} `xml:"evCTeAutorizadoMDFe"`
			} `xml:"detEvento"`
		} `xml:"infEvento"`
	} `xml:"eventoCTe"`
}
