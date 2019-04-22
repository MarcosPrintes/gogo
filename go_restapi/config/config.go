package config

/*
- importando os pacotes necessários para o nosso package
- criando uma struct para passarmos os dados de conexão do nosso db
- estamos lendo o arquivo config.toml; passaremos esses dados para nossa struct
*/

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
