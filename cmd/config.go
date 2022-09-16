package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/redbeestudios/go-seed/pkg"
)

type Config struct {
	Server  pkg.ServerConfig     `json:"server"`
	PokeApi pkg.RestClientConfig `json:"poke_api"`
	Redis   pkg.RedisConfig      `json:"redis"`
}

func InitConfig(env pkg.Env) *Config {
	config := &Config{}

	// TODO: este codigo podria moverse a pkg bajo el nombre de ReadJsonFile,
	// para en el futuro formar parte de una lib de Redbee
	jsonConfig, err := os.Open(fmt.Sprintf("%s_dev.json", env.String()))
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %s", err.Error()))
	}
	defer jsonConfig.Close()

	jsonParser := json.NewDecoder(jsonConfig)
	if err = jsonParser.Decode(config); err != nil {
		panic(fmt.Sprintf("Error parsing config file: %s", err.Error()))
	}

	return config
}
