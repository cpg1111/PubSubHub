package data_store

type Result interface {}

type DataStore interface {
    Get(key string) map[string]string
    Set(key, value string)
}
