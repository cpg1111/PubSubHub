package etcd

import (
    "net/context"

    "github.com/cpg1111/PubSubHub/configReader"

    "github.com/coreos/etcd/client"
)

func New(pSHConf *configReader.Config) *client.KeysApi{
    cfg := etcd.Config{
        Endpoints:               pSHConf.EtcdEndPoints,
        Transport:               client.DefaultTransport,
        HeaderTimeoutPerRequest: time.Second,
    }
    conn, connErr := client.New(cfg)
    if connErr != nil {
        panic(connErr)
    }
    kapi := client.NewKeysAPI(conn)
    return &kapi
}
