package port

type InputDTO struct {
	Message string
}

type KafkaProducer interface {
	Publish(topic string, key, value []byte) error
	Close() error
}
