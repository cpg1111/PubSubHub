package data_store

import (
    "os"
)

type EnvLoader struct {
    env map[string]interface{}
}

func(e *EnvLoader) Set(key, value string){
    // TODO
    // Once PSH master to worker communication is worked out
    // This should tell the master to have all workers set an Env variable to a value
    envErr := os.Setenv(key, value)
    if envErr != nil {
        panic(envErr)
    }
}

func(e .EnvLoader) Get(key string) string{
    var value string
    if e.Env[key] != "" || e.Env[key] != nil {
        value = e.Env[key]
    } else {
        value = os.Getenv(key)
        e.Env[key] = os.Getenv
    }
    return value
}

func NewEnvLoader() *EnvLoader{
    return &EnvLoader{}
}
