package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "log"
    "path"
)

var Config *ConfigFile

type ConfigFile struct {
    Coin, Fiat, Interval string
    Difference int
    Email struct {
        Username, Password, Server string
        Port int
        Recipient string
    }
}

func LoadConfig() {
    file := path.Join(os.Getenv("PWD"), "config.yaml")
    configFile, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(configFile, &Config)
	log.Printf("loaded config: %v", Config)
}