package Controller

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/chapzin/GoSped/Model"
)

func ListaXmls(arquivo string) {
	xmlarquivo, err := os.Open(arquivo)
	if err != nil {
		fmt.Println(err)
	}
	b, _ := ioutil.ReadAll(xmlarquivo)
	isNfePro := string(b[0:70])
	// Estrutura com nfe nao valida
	if strings.Contains(isNfePro, "<NFe xmlns:xsi") {
		var nota2 Model.NFe
		xml.Unmarshal(b, &nota2)
		for _, det := range nota2.InfNFe.Det {
			fmt.Println(det.Prod.CProd)
		}
	}
	// Estrutura com nfe valida
	if strings.Contains(isNfePro, "<nfeProc") {
		var nota Model.NfeProc
		xml.Unmarshal(b, &nota)
		for _, det := range nota.NFe.InfNFe.Det {
			fmt.Println(det.Prod.CProd)
		}
	}

	// Evento nfe inutilizada
	if strings.Contains(isNfePro, "<retInutNFe") {
		var inutilizada Model.RetInutNfe
		xml.Unmarshal(b, &inutilizada)
		fmt.Println(inutilizada)
	}

	// Eventos das Nfe
	if strings.Contains(isNfePro, "<procEventoNFe") {
		var procEventoNfe Model.ProcEventoNFe
		xml.Unmarshal(b, &procEventoNfe)
		fmt.Println(procEventoNfe)
	}

}
