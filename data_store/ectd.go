package data_store

import (
    "net/context"
    "fmt"
    "strings"

    "github.com/cpg1111/PubSubHub/configReader"

    etcd "github.com/coreos/etcd/client"
)

type EtcdStore struct {
    DataStore
    Conn etcd.KeysApi
}

func(e *EtcdStore) SetupConn(pSHConf *configReader.Config){
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
    e.Conn = kapi
}

func(e *EtcdStore) valToString(value string) string{
    newValue := fmt.Sprintf("/%s", strings.ToLower(value))
    return strings.Replace(newValue, "_", "/", nil)
}

func resToMap(res etcd.Nodes, resMap *map[string]interface{}){
    if len(res) > 0 {
        for i:= range res {
            resMap[res[i].Key] = resToMap(res[i].Nodes)
        }
    } else {
        resMap[res.Key] = res.Value
    }
}

func(e *EtcdStore) Get(key string) map[string]interface{} {
    if resMap == nil {
        resMap = make(map[string]interface{})
    }
    res, resErr := e.Conn.Get(context.Background(), e.valToString(key), nil)
    if resErr != nil {
        panic(resErr)
    }
    resToMap(etcd.Nodes{res.Node}, &resMap)
    return resMap
}

func(e *EctdStore) Set(key, value string){
    res, resErr := e.Conn.Set(context.Background(), e.valToString(key), e.valToString(value), nil)
    if resErr != nil {
        panic(nil)
    }
}

func NewEtcd(pSHConf *configReader.Config) *EtcdStore{
    eStore := &EtcdStore{}
    eStore.SetupConn(pSHConf)
    return eStore
}
