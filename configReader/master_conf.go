package configReader

type MasterConf struct {
	defaultConfig
	isWorker   bool
	protocol   string
	backend    string
	pubSubPort int
}

func (master *MasterConf) defaults() {
	if master.isWorker == nil {
		master.isWorker = false
	}
	if master.protocol == nil {
		master.protocol = "ws"
	}
	if master.backend == nil {
		master.backend = "native"
	}
	if master.pubSubPort == nil {
		master.pubSubPort = 8080
	}
}
