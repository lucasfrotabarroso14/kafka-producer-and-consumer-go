package main

import (
	"async_kafka_services/consumer-service/internal/infra/kafka"
	"async_kafka_services/consumer-service/internal/usecase"
)

func main() {

	consumer, err := kafka.NewKafkaConsumer("localhost:9092", "client-consumer-group")
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	exampleUsecase := usecase.NewProcessMessageUseCase()

	err = consumer.Consume("example_topic", func(key []byte, value []byte) {
		if err = exampleUsecase.Execute(nil, key, value); err != nil {
			panic(err)
		}
	})

}
