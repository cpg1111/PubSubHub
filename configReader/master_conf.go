package configReader

type MasterConf struct {
	defaultConfig
	IsWorker   bool
	Protocol   string
	Backend    string
	PubSubPort int
}

func (master *MasterConf) Defaults() {
	if master.IsWorker == nil {
		master.IsWorker = false
	}
	if master.Protocol == nil {
		master.Protocol = "ws"
	}
	if master.Backend == nil {
		master.Backend = "native"
	}
	if master.PubSubPort == nil {
		master.PubSubPort = 8080
	}
}
