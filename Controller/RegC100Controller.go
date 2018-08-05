package Controller

import (
	"github.com/FixAuditoria/GoSped/Model"
)

// RegC100Controller : Responsavel por fazer validacoes no registro C100 do sped
type RegC100Controller struct {
	regC100 Model.RegC100
}

// VerificarChave : Analisa se a nota nao é inutilizada o modelo é 55 e se a chave está vazia e retorna o erro de chave nao informada
func (c *RegC100Controller) VerificarChave(c100 Model.RegC100) string {
	var validacao string
	c.regC100 = c100
	if c.regC100.CodSit != "05" {
		if c.regC100.CodMod == "55" {
			if c.regC100.ChvNfe == "" {
				validacao = "Chave nao informada"
			}
		}
	}
	return validacao
}

// VerificarDtEmissaoChave : Analisa se a nota tem status normal
func (c *RegC100Controller) VerificarDtEmissaoChave(c100 Model.RegC100) string {
	var validacao string
	c.regC100 = c100
	if c.regC100.CodSit == "00" {
		if c.regC100.CodMod == "55" {
			if c.regC100.ChvNfe != "" {
				if len(c.regC100.DtDoc) == 8 {
					chave := c.regC100.ChvNfe
					anoChave := chave[2:4]
					mesChave := chave[4:6]
					dtemit := c.regC100.DtDoc
					anodt := dtemit[6:8]
					mesdt := dtemit[2:4]
					if anoChave != anodt {
						validacao = "Ano Informado de Emissao errado - " + chave
					}
					if mesChave != mesdt {
						validacao = "Mes informado de Emissao errado - " + chave
					}
				}

			}
		}
	}
	return validacao
}
