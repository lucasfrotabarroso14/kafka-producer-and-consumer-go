package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	reader *kafka.Consumer
}

func NewKafkaConsumer(brokers, groupID string) (*Consumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return &Consumer{reader: c}, nil
}

func (c *Consumer) Consume(topic string, groupID string) error {
	err := c.reader.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	for {
		msg, err := c.reader.ReadMessage(-1)
		if err == nil {

			fmt.Printf("Mensagem recebida: %s = %s\n", string(msg.Key), string(msg.Value))
		} else {
			fmt.Printf("Erro: %v\n", err)
		}
	}
}
