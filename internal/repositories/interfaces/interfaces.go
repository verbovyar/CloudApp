package interfaces

type RepoInterface interface {
	Put(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type TransactionLoggerInterface interface {
	WriteDelete(key string)
	WritePut(value, key string)
}
