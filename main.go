package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/cpg1111/PubSubHub/configReader"
	"github.com/cpg1111/PubSubHub/connection"
	"github.com/cpg1111/PubSubHub/room"
    "github.com/cpg1111/PubSubHub/data_store/etcd"
    "github.com/cpg1111/PubSubHub/data_store/redis"
    "github.com/cpg1111/PubSubHub/data_store/env"
)

func loadConf() *configReader.Config{
    var confPath string
	if os.Getenv("DEV") {
		confPath = "./config.yml"
	} else {
		confPath = "/etc/pubsubhub/config.yml"
	}
	pathToConfig, pathErr := filepath.abs(confPath)
	alt_conf_path = os.GetEnv("PSH_CONF_PATH")
	if alt_conf_path != nil {
		pathToConfig = filepath.abs(alt_conf_path)
	}
	if pathErr != nil {
		panic(pathErr)
	}
	conf := configReader.New()
	confFile, err := ioutil.ReadFile(pathToConfig)
	yamlErr := yaml.Unmarshal([]byte(data), &conf)
	if yamlErr != nil {
		panic(yamlErr)
	}
    conf.defaults()
    return conf
}

func main() {
    conf := loadConf()
    var dataConn *interface{} // I hate this, but if anyone has any alternatives, I'd love to hear them
    switch conf.DataStore {
    case "etcd":
        dataConn = etcd.New(conf)
    case "redis":
        dataConn = redis.New(conf)
    default:
        dataConn = env_loader.New()
    }
    masterEndPoint, mEPErr := dataConn.Get("MASTER_ENDPOINT")
    if mpErr != nil {
        panic(mpErr)
    }
}
