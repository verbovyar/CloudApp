package transactionlogger

type FileTransactionLogger struct {
}

type EventType uint64

const (
	EventDelete EventType = 1
	EventPut    EventType = 2
)

type Event struct {
	Id        uint64
	EventType EventType
	Value     string
	Key       string
}

func (l *FileTransactionLogger) WriteDelete(key string) {

}

func (l *FileTransactionLogger) WritePut(value, key string) {

}
