package Controller

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-bongo/bongo"

	"github.com/chapzin/GoSped/Dao"
	"github.com/chapzin/GoSped/Model"
	"github.com/chapzin/GoSped/Utilidades"
)

var path string

func init() {
	path = "empresas"
	pathinvalido := path + "/invalido"
	Utilidades.CriarUmDiretorio(path)
	Utilidades.CriarUmDiretorio(pathinvalido)
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
	tipo := "nfe"
	chave := nota.NFe.InfNFe.Id[3:47]
	cnpj := nota.NFe.InfNFe.Emit.Cnpj
	ano := "20" + nota.NFe.InfNFe.Id[5:7]
	mes := nota.NFe.InfNFe.Id[7:9]
	pathFinal := Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = pathFinal + "/" + chave + ".xml"
	return pathArquivo, nota
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

// ProcessarXmls : Funcao responsavel por verificar o que contem e mandar para o destino correto
func ProcessarXmls(arquivo string, conn *bongo.Connection) {
	xmlByte := TransformarXmlEmByte(arquivo)
	isNfePro := string(xmlByte[0:70])

	if strings.Contains(isNfePro, "><NFe xmlns=") {
		pathArquivo, nota := ProcessarNfeSemProcValido(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirNfeProcSemValidade(Dao.COLLECTIONNFESEMVALIDADE, nota, conn)
	}

	if strings.Contains(isNfePro, "<NFe xmlns:xsi") {
		pathArquivo, nota := ProcessarNfeSemValidade(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirNfeProcSemValidade(Dao.COLLECTIONNFESEMVALIDADE, nota, conn)
	}
	// Estrutura com nfe valida
	if strings.Contains(isNfePro, "<nfeProc") {
		pathArquivo, nota := ProcessarNfeProcValido(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirNfeProc(Dao.COLLECTIONNFEPROC, nota, conn)
	}

	// Evento nfe inutilizada
	if strings.Contains(isNfePro, "<retInutNFe") {
		pathArquivo, inutilizada := ProcessarRetornoNfeInutilizada(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirRetornoNfeInutilizada(Dao.COLLECTIONNFEINUTILIZADA, inutilizada, conn)
	}

	// Eventos das Nfe
	if strings.Contains(isNfePro, "<procEventoNFe") {
		pathArquivo, procEventoNfe := ProcessarEventoNfe(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirEventoNfe(Dao.COLLECTIONNFEEVENTO, procEventoNfe, conn)
	}

	if strings.Contains(isNfePro, "<procEventoCTe versao") {
		pathArquivo, procEventoCte := ProcessarEventoCte(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirEventoCte(Dao.COLLECTIONCTEEVENTO, procEventoCte, conn)
	}

}
