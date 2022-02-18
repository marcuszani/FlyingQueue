package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var path = "config.yml"

type config struct {
	Name string `yaml:"Name"`
	Url  string `yaml:"URL"`
	Port string `yaml:"Port"`
}

var Cfg *config = &config{}

func Configuration() {
	checkConfigFile()

}

func init() {

	//Config := &config{}

	//bs, err := os.ReadFile("config.yml")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = yaml.Unmarshal(bs, Config)

	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func checkConfigFile() bool {

	fmt.Printf("Verificando Arquivo de configuração: ")

	_, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("não existe")
		err = createConfigFile()

		if err != nil {
			fmt.Println(err)
		}

		return false
	} else {
		fmt.Println("OK")
		loadConfigFile()
		return true
	}
}

func createConfigFile() error {
	bs, _ := yaml.Marshal(Cfg)
	return os.WriteFile(path, bs, 0600)
}

func loadConfigFile() {
	Cfg := &config{}

	bs, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(bs, Cfg)

	if err != nil {
		fmt.Println(err)
	}
}
