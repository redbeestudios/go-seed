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
}

func InitConfig(env pkg.Env) *Config {
	config := &Config{}

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
