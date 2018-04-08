package Controller

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-bongo/bongo"

	"github.com/chapzin/GoSped/Model"
)

var path string

func init() {
	path = "empresas"
	pathinvalido := path + "/invalido"
	IsCreateDir(path)
	IsCreateDir(pathinvalido)
}

func IsCreateDir(path string) {
	// Caminho do cnpj da empresa emitente
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

func CriarPastas(path string, cnpj string, ano string, mes string, tipo string) {
	pathcnpj := path + "/" + cnpj
	fmt.Println(pathcnpj)
	pathano := path + "/" + cnpj + "/" + ano
	pathmes := path + "/" + cnpj + "/" + ano + "/" + mes
	pathtipo := path + "/" + cnpj + "/" + ano + "/" + mes + "/" + tipo

	IsCreateDir(pathcnpj)
	IsCreateDir(pathano)
	IsCreateDir(pathmes)
	IsCreateDir(pathtipo)
}

func MoveXml(pathInicial string, pathFinal string) {
	err := os.Rename(pathInicial, pathFinal)
	if err != nil {
		fmt.Println(err)
	}
}

func ListaXmls(arquivo string, conn *bongo.Connection) {
	xmlarquivo, err := os.Open(arquivo)
	if err != nil {
		fmt.Println(err)
	}
	b, _ := ioutil.ReadAll(xmlarquivo)
	xmlarquivo.Close()
	isNfePro := string(b[0:70])
	// fmt.Println(isNfePro)
	// Estrutura com nfe nao valida

	if strings.Contains(isNfePro, "><NFe xmlns=") {
		var nota Model.NFe
		xml.Unmarshal(b, &nota)
		tipo := "nfe"
		chave := nota.InfNFe.Id[3:47]
		cnpj := nota.InfNFe.Emit.Cnpj
		ano := "20" + nota.InfNFe.Id[5:7]
		mes := nota.InfNFe.Id[7:9]
		pathArquivo := path + "/" + cnpj + "/" + ano + "/" + mes + "/" + tipo + "/" + chave + ".xml"
		CriarPastas(path, cnpj, ano, mes, tipo)
		MoveXml(arquivo, pathArquivo)

	}

	if strings.Contains(isNfePro, "<NFe xmlns:xsi") {
		var nota2 Model.NFe
		xml.Unmarshal(b, &nota2)
		tipo := "nfeSemValidade"
		chave := nota2.InfNFe.Id[3:47]
		cnpj := nota2.InfNFe.Emit.Cnpj
		ano := "20" + nota2.InfNFe.Id[5:7]
		mes := nota2.InfNFe.Id[7:9]
		pathArquivo := path + "/" + cnpj + "/" + ano + "/" + mes + "/" + tipo + "/" + chave + ".xml"
		CriarPastas(path, cnpj, ano, mes, tipo)
		MoveXml(arquivo, pathArquivo)
		// for _, det := range nota2.InfNFe.Det {
		// 	fmt.Println(det.Prod.CProd)
		// }
	}
	// Estrutura com nfe valida
	if strings.Contains(isNfePro, "<nfeProc") {
		var nota Model.NfeProc
		xml.Unmarshal(b, &nota)
		tipo := "nfe"
		chave := nota.NFe.InfNFe.Id[3:47]
		cnpj := nota.NFe.InfNFe.Emit.Cnpj
		ano := "20" + nota.NFe.InfNFe.Id[5:7]
		mes := nota.NFe.InfNFe.Id[7:9]
		pathArquivo := path + "/" + cnpj + "/" + ano + "/" + mes + "/" + tipo + "/" + chave + ".xml"
		CriarPastas(path, cnpj, ano, mes, tipo)
		MoveXml(arquivo, pathArquivo)
		err := conn.Collection("nfeProc").Save(&nota)
		if err != nil {
			fmt.Println(err)
		}
		// for _, det := range nota.NFe.InfNFe.Det {
		// 	fmt.Println(det.Prod.CProd)
		// }
	}

	// Evento nfe inutilizada
	if strings.Contains(isNfePro, "<retInutNFe") {
		var inutilizada Model.RetInutNfe
		xml.Unmarshal(b, &inutilizada)
		cnpj := inutilizada.InfInut.CNPJ
		ano := "20" + inutilizada.InfInut.Ano
		nfini := inutilizada.InfInut.NNFIni
		nffin := inutilizada.InfInut.NNFFin
		tipo := "nfeInutilizadas"
		pathInut := path + "/" + cnpj + "/" + ano + "/" + tipo
		pathArquivo := pathInut + "/" + nfini + "-" + nffin + ".xml"

		IsCreateDir(pathInut)
		MoveXml(arquivo, pathArquivo)
	}

	// Eventos das Nfe
	if strings.Contains(isNfePro, "<procEventoNFe") {
		var procEventoNfe Model.ProcEventoNFe
		xml.Unmarshal(b, &procEventoNfe)
		cnpj := procEventoNfe.Evento.InfEvento.CNPJ
		tipo := "evento"
		chave := procEventoNfe.Evento.InfEvento.ChNFe
		ano := "20" + chave[2:4]
		mes := chave[4:6]
		pathArquivo := path + "/" + cnpj + "/" + ano + "/" + mes + "/" + tipo + "/procEvent-" + chave + ".xml"
		CriarPastas(path, cnpj, ano, mes, tipo)
		MoveXml(arquivo, pathArquivo)
	}

}
