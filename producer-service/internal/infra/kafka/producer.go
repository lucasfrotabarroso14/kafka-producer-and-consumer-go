package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducerImpl struct {
	producer *kafka.Producer
}

func NewKafkaProducer(brokers string) (*KafkaProducerImpl, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {

					println("Erro ao enviar:", ev.TopicPartition.Error.Error())
				} else {

					println("Mensagem entregue em:", ev.TopicPartition.String())
				}
			}
		}
	}()

	return &KafkaProducerImpl{producer: p}, nil
}

func (kp *KafkaProducerImpl) Publish(topic string, key, value []byte) error {
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value,
	}, nil)
}

func (kp *KafkaProducerImpl) Close() error {
	kp.producer.Flush(15000)
	kp.producer.Close()
	return nil
}
