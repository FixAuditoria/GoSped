package Controller

import (
	"os"
	"strings"

	"github.com/go-bongo/bongo"

	"fmt"
	"github.com/chapzin/GoSped/Dao"
	"github.com/chapzin/GoSped/Utilidades"
)

var path string

func init() {
	path = "empresas"
	pathinvalido := path + "/invalido"
	Utilidades.CriarUmDiretorio(path)
	Utilidades.CriarUmDiretorio(pathinvalido)
}

// ProcessarXmls : Funcao responsavel por verificar o que contem e mandar para o destino correto
func ProcessarXmls(arquivo string, conn *bongo.Connection) {
	xmlByte := TransformarXmlEmByte(arquivo)
	isNfePro := string(xmlByte[0:120])
	fmt.Println(isNfePro)
	//fmt.Println(arquivo)
	//panic(isNfePro)
	if strings.Contains(isNfePro, "><NFe xmlns=") && !strings.Contains(isNfePro, "<NFeDFe") {
		pathArquivo, nota := ProcessarNfeSemProcValido(xmlByte)
		if pathArquivo != "" {
			Utilidades.MoverArquivos(arquivo, pathArquivo)
			Dao.InserirNfeProcSemValidade(Dao.COLLECTIONNFESEMVALIDADE, nota, conn)
		}
	}

	if strings.Contains(isNfePro, "<NFe xmlns:xsi") {
		pathArquivo, nota := ProcessarNfeSemValidade(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirNfeProcSemValidade(Dao.COLLECTIONNFESEMVALIDADE, nota, conn)
	}
	// Estrutura com nfe valida
	if strings.Contains(isNfePro, "<nfeProc") {
		pathArquivo, nota := ProcessarNfeProcValido(xmlByte)
		// Verifica se arquivo nao foi realmente preenchido
		if pathArquivo != "" {
			Utilidades.MoverArquivos(arquivo, pathArquivo)
			Dao.InserirNfeProc(Dao.COLLECTIONNFEPROC, nota, conn)
		}

	}

	if strings.Contains(isNfePro, "<nfeProc versao") && strings.Contains(isNfePro, "<nfeProc xmlns") {
		pathArquivo, nota2 := ProcessarNfeProcNfeProc(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirNfeProc(Dao.COLLECTIONNFEPROC, nota2, conn)
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

	if strings.Contains(isNfePro, "<procEventoCTe") {
		pathArquivo, procEventoCte := ProcessarEventoCte(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirEventoCte(Dao.COLLECTIONCTEEVENTO, procEventoCte, conn)
	}

	if strings.Contains(isNfePro, "<cteProc") {
		pathArquivo, procCte := ProcessarCteProc(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirCteProc(Dao.COLLECTIONCTE, procCte, conn)
	}

	if strings.Contains(isNfePro, "<outrasInfDANFe") {
		os.Remove(arquivo)
	}

	if strings.Contains(isNfePro, "<retConsReciNFe") {
		os.Remove(arquivo)
	}

	if strings.Contains(isNfePro, "<resEvento") {
		pathArquivo, resEvento := ProcessarResEvento(xmlByte)
		Utilidades.MoverArquivos(arquivo, pathArquivo)
		Dao.InserirResEvento(Dao.COLLECTIONRESEVENTO, resEvento, conn)
	}

	if strings.Contains(isNfePro, "<NFeDFe") {
		os.Remove(arquivo)
	}

	if strings.Contains(isNfePro, "<infEvento") {
		os.Remove(arquivo)
	}

}
