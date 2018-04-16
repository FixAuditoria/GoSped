package Controller

import (
	"encoding/xml"
	"fmt"
	"github.com/chapzin/GoSped/Model"
	"github.com/chapzin/GoSped/Utilidades"
	"io/ioutil"
	"os"
)

func DeleteArquivoVazio(arquivo string) {
	file, err := os.Open(arquivo)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}
	if stat.Size() <= 0 {
		file.Close()
		os.Remove(arquivo)
	}
}

// ProcessarNfeSemProcValido : Responsavel por executar estrutura de pastas e retornar o novo path do arquivo
func ProcessarNfeSemProcValido(file []byte) (pathArquivo string, nota Model.NFe) {
	xml.Unmarshal(file, &nota)
	tipo := "nfe"
	chave := nota.InfNFe.Id[3:47]
	cnpj := nota.InfNFe.Emit.Cnpj
	ano := "20" + nota.InfNFe.Id[5:7]
	mes := nota.InfNFe.Id[7:9]
	pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathFinal + "/" + chave + ".xml"

	return pathArquivo, nota
}

func ProcessarResEvento(file []byte) (pathArquivo string, resEvento Model.ResEvento) {
	xml.Unmarshal(file, &resEvento)
	tipo := "resEvento"
	chave := resEvento.ChNFe
	cnpj := resEvento.CNPJ
	ano := resEvento.DhEvento[0:4]
	mes := resEvento.DhEvento[5:7]
	pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathFinal + "/" + chave + ".xml"
	return pathArquivo, resEvento
}

func ProcessarCteProc(file []byte) (pathArquivo string, cte Model.CteProc) {
	xml.Unmarshal(file, &cte)
	tipo := "cte"
	chave := cte.Cte.InfCte.Id[3:47]
	cnpj := cte.Cte.InfCte.Dest.CNPJ
	data := cte.Cte.InfCte.Ide.DhEmi
	ano := data[0:4]
	mes := data[5:7]
	pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathFinal + "/" + chave + ".xml"
	return pathArquivo, cte

}

// ProcessarNfeSemValidade : Responsavel por executar estrutura de pastas dos xmls sem validade e retorna o path do arquivo
func ProcessarNfeSemValidade(file []byte) (pathArquivo string, nota2 Model.NFe) {
	xml.Unmarshal(file, &nota2)
	tipo := "nfeSemValidade"
	chave := nota2.InfNFe.Id[3:47]
	cnpj := nota2.InfNFe.Emit.Cnpj
	ano := "20" + nota2.InfNFe.Id[5:7]
	mes := nota2.InfNFe.Id[7:9]
	pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathFinal + "/" + chave + ".xml"

	return pathArquivo, nota2
}

// ProcessarNfeProcValido : Funcao responsavel separar variaveis e retornar path e model
func ProcessarNfeProcValido(file []byte) (pathArquivo string, nota Model.NfeProc) {
	xml.Unmarshal(file, &nota)

	if len(nota.NFe.InfNFe.Id) <= 0 {
		return "", nota
	} else {
		tipo := "nfe"
		chave := nota.NFe.InfNFe.Id[3:47]
		cnpj := nota.NFe.InfNFe.Emit.Cnpj
		ano := "20" + nota.NFe.InfNFe.Id[5:7]
		mes := nota.NFe.InfNFe.Id[7:9]
		pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
		pathArquivo = pathFinal + "/" + chave + ".xml"
		return pathArquivo, nota
	}

}

func ProcessarNfeProcNfeProc(file []byte) (pathArquivo string, nota Model.NfeProc) {
	var nota2 Model.NfeProcNfeProc
	xml.Unmarshal(file, &nota2)
	tipo := "nfe"
	chave := nota2.NfeProc.NFe.InfNFe.Id[3:47]
	cnpj := nota2.NfeProc.NFe.InfNFe.Emit.Cnpj
	ano := "20" + nota2.NfeProc.NFe.InfNFe.Id[5:7]
	mes := nota2.NfeProc.NFe.InfNFe.Id[7:9]
	pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathFinal + "/" + chave + ".xml"
	return pathArquivo, nota2.NfeProc
}

// TransformarXmlEmByte : Funcao responsavel por receber path do arquivo xml e retornar em byte
func TransformarXmlEmByte(arquivo string) (b []byte) {
	xmlarquivo, err := os.Open(arquivo)
	if err != nil {
		fmt.Println(err)
	}
	b, _ = ioutil.ReadAll(xmlarquivo)
	xmlarquivo.Close()
	return b
}

// ProcessarRetornoNfeInutilizada : Funcao responsavel por analisar xml de retorno inutilizado e retornar o path detinado
func ProcessarRetornoNfeInutilizada(file []byte) (pathArquivo string, inutilizada Model.RetInutNfe) {
	xml.Unmarshal(file, &inutilizada)
	cnpj := inutilizada.InfInut.CNPJ
	ano := "20" + inutilizada.InfInut.Ano
	nfini := inutilizada.InfInut.NNFIni
	nffin := inutilizada.InfInut.NNFFin
	tipo := "nfeInutilizadas"
	pathInut := path + "/" + cnpj + "/" + ano + "/" + tipo
	pathArquivo = pathInut + "/" + nfini + "-" + nffin + ".xml"
	Utilidades.CriarUmDiretorio(pathInut)
	return pathArquivo, inutilizada
}

// ProcessarEventoNfe : Funcao responsavel por analisar xml de evento e retornar o path destinado
func ProcessarEventoNfe(file []byte) (pathArquivo string, procEventoNfe Model.ProcEventoNFe) {
	xml.Unmarshal(file, &procEventoNfe)
	cnpj := procEventoNfe.Evento.InfEvento.CNPJ
	tipo := "evento"
	chave := procEventoNfe.Evento.InfEvento.ChNFe
	ano := "20" + chave[2:4]
	mes := chave[4:6]
	tpevento := procEventoNfe.Evento.InfEvento.TpEvento
	pathfinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathfinal + "/procEvent-" + tpevento + "-" + chave + ".xml"
	return pathArquivo, procEventoNfe
}

func ProcessarEventoCte(file []byte) (pathArquivo string, procEventoCte Model.ProcEventoCTe) {
	xml.Unmarshal(file, &procEventoCte)
	cnpj := procEventoCte.EventoCTe.InfEvento.DetEvento.EvCTeAutorizadoMDFe.Emit.CNPJ
	ano := procEventoCte.EventoCTe.InfEvento.DhEvento[0:4]
	mes := procEventoCte.EventoCTe.InfEvento.DhEvento[5:7]
	chave := procEventoCte.EventoCTe.InfEvento.ChCTe
	tpEvento := procEventoCte.EventoCTe.InfEvento.TpEvento
	tipo := "eventoCte"
	pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathFinal + "/" + chave + "-" + tpEvento + ".xml"
	return pathArquivo, procEventoCte
}
