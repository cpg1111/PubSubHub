package main

import(
	"log"
	"net"
	"os"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"configReader"
	"connection"
	"room"
)

func main() {
	var confPath string
	if os.GetEnv('DEV') {
		confPath = "./config.yml"
	} else {
		confPath = "/etc/pubsubhub/config.yml"
	}
	PATH_TO_CONFIG, pathErr := filepath.abs(confPath)
	alt_conf_path = os.GetEnv("PSH_CONF_PATH")
	if alt_conf_path != nil {
		PATH_TO_CONFIG = filepath.abs(alt_conf_path)
	}
	if pathErr != nil {
		panic(pathErr)
	}
	conf := config{}
	conf.defaults()
	confFile, err := ioutil.ReadFile(PATH_TO_CONFIG)
	yamlErr := yaml.Unmarshal([]byte(data), &t)
	if yamlErr != nil {
		log.Fatal(yamlErr)
	}

}
