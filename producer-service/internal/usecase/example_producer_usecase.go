package usecase

import (
	"async_kafka_services/producer-service/internal/domain/port"
	"context"
)

type ExampleProducerUseCase struct {
	kafka port.KafkaProducer
}

func NewExampleProducerUseCase(kafka port.KafkaProducer) *ExampleProducerUseCase {
	return &ExampleProducerUseCase{
		kafka: kafka,
	}
}

func (uc *ExampleProducerUseCase) Execute(ctx context.Context, input port.InputDTO) (int, error) {
	err := uc.kafka.Publish("example_topic", []byte("key"), []byte(input.Message))
	if err != nil {
		return 500, err
	}
	return 200, nil
}
