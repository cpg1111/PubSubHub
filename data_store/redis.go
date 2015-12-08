package data_store

import (
    "github.com/cpg1111/PubSubHub/configReader"

    redis "github.com/xuyu/goredis"
)

type Redis struct {
    DataStore
    Conn *redis.Redis
}

func(r *Redis) SetupConn(conf *configReader.Config){
    conn, connErr := redis.Dial(&redis.DialConfig{Address: conf.RedisAddress})
    if connErr != nil {
        panic(connErr)
    }
    r.Conn = conn
}

func(r *Redis) Get(key string) string{
    res, resErr := r.Conn.Get(key)
    if resErr != nil {
        panic(nil)
    }
    return string(res)
}

func(r *Redis) Set(key, value string){
    res, resErr := r.Conn.GetSet(key, value)
    if resErr != nil {
        panic(resErr)
    }
}

func NewRedis() *Redis{
    return &Redis{}
}
