package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "log"
    "path"
)

//Global configuration containing SMTP auth, prefs, etc.
var Config *ConfigFile

//YAML structure of the config.yaml file.
type ConfigFile struct {
    Coin, Fiat, Interval string
    Difference float64
    Email struct {
        Username, Password, Server string
        Port int
        Recipient string
    }
}

//LoadConfig loads the config from the config.yaml
//file to the global variable Config as type *ConfigFile.
func LoadConfig() {
    file := path.Join(os.Getenv("PWD"), "config.yaml")
    configFile, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(configFile, &Config)
	log.Printf("loaded config: %v", Config)
}