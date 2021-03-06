package models

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MailServer struct {
	PrimaryDomain            string `yaml:"PrimaryDomain"`
	GlobalDomainMaxUserCount int    `yaml:"GlobalDomainMaxUserCount"`
	GlobalDomainMaxUserQuota int    `yaml:"GlobalDomainMaxUserQuota"`
	GlobalDomainMaxMailCount int    `yaml:"GlobalDomainMaxMailCount"`
}

// WebServer
type WebServer struct {
	Listen string `yaml:"Listen"`
}

// MySQL
type MySQL struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
}

// ServerConfig
type ServerConfig struct {
	MailServer MailServer `yaml:"MailServer"`
	WebServer  WebServer  `yaml:"WebServer"`
	MySQL      MySQL      `yaml:"MySQL"`
}

// InitConfig
func InitConfig() (cfg ServerConfig) {
	configData, err := ioutil.ReadFile("partner.yaml")
	if err != nil {
		log.Fatal("partner.yaml not found!")
	}
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		panic(err)
	}
	return
}
