package ConfigTom

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config : Usando sistema de TOML para configuracoes de conexao de banco de dados mongo
type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
