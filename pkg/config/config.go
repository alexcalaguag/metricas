package config

import (
	"os"
	"log"
	"github.com/ilyakaznacheev/cleanenv"
)

type SonarConfiguration struct {
	User struct {
		Chave   string `yaml:"chave"`
		Senha   string `yaml:"senha"`
	} `yaml:"user"`
	Database struct {
		URL     string `yaml:"url"`
		DB      string `yaml:"db"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"database"`
	Sonar struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		SufixTestCoverage string `yaml:"sufixTestCoverage"`
	} `yaml:"server"`
}


var Cfg SonarConfiguration


func InitConfig() {
	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig("../pkg/config/config.yaml", &Cfg); err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	log.Print("Configuration loaded successfully")
}
