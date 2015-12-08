package data_store

type DataStore interface {
    Get(key string) string
    Set(key, value string)
}
