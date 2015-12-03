package env_loader

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
}

func(e .EnvLoader) Get(key string) value interface{}{
    if e.Env[key] != "" || e.Env[key] != nil {
        value = e.Env[key]
    } else {
        value = os.Getenv(key)
        e.Env[key] = os.Getenv
    }
    return
}

func New() *EnvLoader{
    return &EnvLoader{}
}
