"# GoSped" 
# Analise e Auditoria EFD Fiscal em Go
[![Build Status](https://travis-ci.org/chapzin/GoSped.svg?branch=master)](https://travis-ci.org/chapzin/GoSped)
[![Go Report Card](https://goreportcard.com/badge/github.com/chapzin/GoSped)](https://goreportcard.com/report/github.com/chapzin/GoSped)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)
[![Join the chat at https://gitter.im/olebedev/go-starter-kit](https://img.shields.io/gitter/room/nwjs/nw.js.svg?maxAge=2592000&style=plastic)](https://gitter.im/GoAuditoriaFiscal/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)
[![Donate](https://img.shields.io/badge/Donate-PayPal-blue.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=R673QGW2LQVCJ)


Projeto voltado para fazer analise, armazenamento e correções de sped fiscal e relatorios que ajudam na auditoria fiscal de uma empresa.

![Sped](sped-maior.png "Sped")
## O que é o Sped Fiscal?
A Escrituração Fiscal Digital - EFD é um arquivo digital, que se constitui de um conjunto de escriturações de documentos fiscais e de outras informações de interesse dos Fiscos das unidades federadas e da Secretaria da Receita Federal do Brasil, bem como de registros de apuração de impostos referentes às operações e prestações praticadas pelo contribuinte.
Este arquivo deverá ser assinado digitalmente e transmitido, via Internet, ao ambiente Sped.

###### SITE OFICIAL DO SPED: http://sped.rfb.gov.br/

## Como compilar 
```
go get github.com/chapzin/GoSped
go build
```

## Como utilizar
- Adicione todos xmls e speds do periodo onde pretende fazer a importação na pasta speds

## Funcionalidades que serão desenvolvidas no sistema:
- Importar todos Speds e Xmls(Cte, nfe) e organizar na pasta empresas pelo cnpj do emitente;
- Acessar via web o sistema onde vao está todos xmls, analise e validacoes dos speds, processamentos de estoques;
- Enviar relatórios e arquivos dos novos inventários por email;

## Motivação
Não é segredo que no mundo Fiscal existe uma enorme complexidade, com mudanças e atualizações na legislação. O calendário fiscal das empresas é lotado de obrigações a serem cumpridas, com datas definidas pelo FISCO.

Muitas vezes, as obrigações são entregues com a premissa de cumprimento de prazo e a qualidade da informação acaba ficando em segundo plano. Dessa forma, as necessidades das retificações se acumulam e faz com que o mês inicie com “dias faltando”.

Além de ter o importantíssimo papel de atender todas as exigências estabelecidas pelo FISCO Municipal, Estadual e Federal, muitas vezes o setor fiscal das empresas acaba tendo que corrigir tudo o que saiu ou entrou errado durante o mês.

De acordo com a experiencia que acumulei em auditoria fiscal nesses ultimos 5 anos quero desenvolver um sistema para facilitar a vida de um contador, auditor e de empresas que queiram armazenar e analisar seus arquivos fiscais, espero que com esse sistema as empresas e contadores consigam encontrar com mais facilidade os erros apresentados e corrigir assim evitando multas altissimas aplicadas pelo fisco, acredito que as empresas devam pagar realmente apenas pelo que está errado e não pela falta de conhecimento ou até mesmo ma fé de fiscais em relação ao conteúdo informado nos speds.

O principal objetivo do sistema é facilitar que as empresas possam fazer um analise da qualidade das informações enviadas, com isso podendo fazer retificações de forma voluntária evitando receber uma multa muito alta.

## Dúvidas?

Abra um issue na página do projeto no GitHub ou [clique aqui](https://github.com/chapzin/parse-efd-fiscal/issues).

## Donate
Ajude a acabar com as injustiça feita pela SEFAZ devido a tantas obrigações a serem entregues.

Donate via [PayPal](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=R673QGW2LQVCJ)

## Colaboradores

- Ricardo Gomes (https://github.com/chapzin)
- Cesar Gimenes (https://github.com/crgimenes)
- Rafael Marques(https://github.com/RafaMarquesLearn)

## License

The project Go Auditoria Fiscal is available under the [MIT license](LICENSE).
