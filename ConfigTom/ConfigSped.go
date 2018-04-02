package ConfigTom

import (
	"log"

	"github.com/BurntSushi/toml"
)

// ConfigSped : Usando sistema de TOML para configuracoes de conexao de banco de dados mongo
type ConfigSped struct {
	Server      string
	Database    string
	DataBaseNfe string
	PathImport  string
}

func (c *ConfigSped) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
