package Controller

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chapzin/GoSped/Model"
)

func ListaXmls(arquivo string) {
	xmlarquivo, err := os.Open(arquivo)
	if err != nil {
		fmt.Println(err)
	}
	var nota Model.NfeProc
	b, _ := ioutil.ReadAll(xmlarquivo)
	xml.Unmarshal(b, &nota)
	if nota.NFe.InfNFe.Versao != "3.10" {
		var nota2 Model.NFe
		xml.Unmarshal(b, &nota2)
		//fmt.Println(nota2.InfNFe.Det.Prod.CProd)
		for _, det := range nota2.InfNFe.Det {
			fmt.Println(det.Prod.CProd)
		}
	} else {
		for _, det := range nota.NFe.InfNFe.Det {
			fmt.Println(det.Prod.CProd)
		}
	}

}
