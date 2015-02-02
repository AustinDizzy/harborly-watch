package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

//Global configuration containing SMTP auth, prefs, etc.
var Config *ConfigFile

//ConfigFile is YAML structure of config.yaml file.
type ConfigFile struct {
	Coin, Fiat, Interval string
	Difference           float64
	Email                struct {
		Username, Password, Server string
		Port                       int
		Recipient                  string
	}
}

//LoadConfig loads the config from the config.yaml
//file to the global variable Config as type *ConfigFile.
func LoadConfig() {
	file := path.Join(GetDir(), "config.yaml")
	configFile, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(configFile, &Config)
	if Config != nil {
		log.Printf("loaded config: %v", Config)
	} else {
		log.Printf("Error loading config.yaml file in %v.", GetDir())
		os.Exit(1)
	}
}

//GetDir return the current directory holding the executable
func GetDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
