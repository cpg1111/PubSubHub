package redis

import (
    "github.com/cpg1111/PubSubHub/configReader"

    redis "github.com/xuyu/goredis"
)

func New(conf *configReader.Config) *redis.Redis{
    conn, connErr := redis.Dial(&redis.DialConfig{Address: conf.RedisAddress})
    if connErr != nil {
        panic(connErr)
    }
    return conn
}
