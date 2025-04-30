package port

type KafkaConsumer interface {
	Consume(topic string, handler func(key []byte, value []byte)) error
	Close() error
}
