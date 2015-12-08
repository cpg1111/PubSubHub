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
    "github.com/cpg1111/PubSubHub/data_store"
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

func setupMasterConn(isMaster bool, masterEndPoint string) connection.Connection, *map[string]room.Channel{
    var rooms map[string]room.Channel
    masterConn := connection.NewTCP()
    if isMaster {
        rooms = make(map[string]room.Channel)
        masterConn.Serve(masterEndPoint, &rooms)
    }
    masterConn.Connect(masterEndPoint)
    if !isMaster {
        masterConn.GetRooms()
    }
}

func main() {
    hostIp := net.InterfaceAddrs()[1]
    conf := loadConf()
    var dataConn *data_store.DataStore
    switch conf.DataStore {
    case "etcd":
        dataConn = data_store.NewEtcd(conf)
    case "redis":
        dataConn = data_store.NewRedis(conf)
    default:
        dataConn = data_store.NewEnvLoader()
    }
    masterEndPoint, mEPErr := dataConn.Get("MASTER_ENDPOINT")
    if mpErr != nil {
        panic(mpErr)
    }
    hostEndPoint := net.JoinHostPort(hostIp, string(conf.MasterPort))
    if masterEndPoint == nil {
        dataConn.Set("MASTER_ENDPOINT", hostEndPoint)
        masterEndPoint = dataConn.Get("MASTER_ENDPOINT")
    }
    isMaster := (hostEndPoint == masterEndPoint)
    setupMasterConn(isMaster, dataConn)
}
