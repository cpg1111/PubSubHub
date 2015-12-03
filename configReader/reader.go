package configReader

import (
	"gopkg.in/yaml.v2"
)

type defaultConfig interface {
	Defaults()
}

type Config struct {
	defaultConfig
	Master     MasterConf
	MinWorkers int
	MaxWorkers int
    DataStore string
    EtcdEndPoints []string
    RedisAddress string
}

func (conf *config) Defaults() {
	conf.Master.Defaults()
	if conf.MinWorkers == nil {
		conf.MinWorkers = 1
	}
	if conf.MaxWorkers == nil {
		conf.MaxWorkers = 5
	}
}

func New() *Config{
    return &Config{}
}
