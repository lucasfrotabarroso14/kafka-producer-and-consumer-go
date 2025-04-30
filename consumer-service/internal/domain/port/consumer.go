package port

type KafkaConsumer interface {
	Consume(topic string, groupID string) error
	Close() error
}
