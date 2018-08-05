package Dao

import (
	"fmt"

	"github.com/FixAuditoria/GoSped/Model"
	"github.com/go-bongo/bongo"
)

// InserirNfeProcMongo : Inserir no banco de dados mongo com o bongo
func InserirNfeProc(colecao string, nota Model.NfeProc, conn *bongo.Connection) {
	err := conn.Collection(colecao).Save(&nota)
	if err != nil {
		fmt.Println(err)
	}
}

func InserirNfeProcSemValidade(colecao string, nota Model.NFe, conn *bongo.Connection) {
	err := conn.Collection(colecao).Save(&nota)
	if err != nil {
		fmt.Println(err)
	}
}

func InserirRetornoNfeInutilizada(colecao string, inutilizada Model.RetInutNfe, conn *bongo.Connection) {
	err := conn.Collection(colecao).Save(&inutilizada)
	if err != nil {
		fmt.Println(err)
	}
}

func InserirEventoNfe(colecao string, procEventoNfe Model.ProcEventoNFe, conn *bongo.Connection) {
	err := conn.Collection(colecao).Save(&procEventoNfe)
	if err != nil {
		fmt.Println(err)
	}
}

func InserirEventoCte(colecao string, procEventoCte Model.ProcEventoCTe, conn *bongo.Connection) {
	err := conn.Collection(colecao).Save(&procEventoCte)
	if err != nil {
		fmt.Println(err)
	}
}

func InserirCteProc(colecao string, procCte Model.CteProc, conn *bongo.Connection) {
	err := conn.Collection(colecao).Save(&procCte)
	if err != nil {
		fmt.Println(err)
	}
}

func InserirResEvento(colecao string, resEvento Model.ResEvento, conn *bongo.Connection) {
	err := conn.Collection(colecao).Save(&resEvento)
	if err != nil {
		fmt.Println(err)
	}
}
