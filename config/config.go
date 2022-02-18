package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var path = "config.yml"

type config struct {
	Server   map[string]string `json:"Server" yaml:"Server"`
	Database map[string]string `json:"Database" yaml:"Database"`
	// Name   string `json:"Name" yaml:"Name"`
	// Url    string `json:"Url"  yaml:"URL"`
	// Port   int    `json:"Port" yaml:"Port"`
}

var Cfg *config = &config{}

func init() {
	fmt.Printf("Verificando Arquivo de configuração: ")

	_, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("não existe")
		err = createConfigFile()
		loadConfigFile()

		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("OK")
		loadConfigFile()
	}
}

func createConfigFile() error {
	// Cfg := &config{
	// 	Name: "DT Solucoes",
	// 	Url:  "localhost",
	// 	Port: 8080,
	// }

	Cfg.Server = map[string]string{
		"Name": "DT Solucoes",
		"Url":  "localhost",
		"Port": "8080",
	}
	Cfg.Database = map[string]string{
		"Address":    "localhost",
		"Port":       "5432",
		"DbName":     "flyingqueue",
		"DbUserName": "root",
		"DbPassword": "root",
	}

	bs, _ := yaml.Marshal(Cfg)
	return os.WriteFile(path, bs, 0600)
}

func loadConfigFile() {

	bs, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(bs, Cfg)

	for key, value := range Cfg.Server {
		os.Setenv(key, value)
	}

	if err != nil {
		fmt.Println(err)
	}
}
