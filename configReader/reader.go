package configReader

import (
	"gopkg.in/yaml.v2"
)

type defaultConfig interface {
	defaults()
}

type config struct {
	defaultConfig
	master     MasterConf
	minWorkers int
	maxWorkers int
}

func (conf *config) defaults() {
	conf.master.defaults()
	if conf.minWorkers == nil {
		conf.minWorkers = 1
	}
	if conf.maxWorkers == nil {
		conf.maxWorkers = 5
	}
}
