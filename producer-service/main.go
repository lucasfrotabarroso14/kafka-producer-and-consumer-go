package main

import (
	"async_kafka_services/producer-service/internal/infra/kafka"
	"async_kafka_services/producer-service/internal/infra/webserver"
	"async_kafka_services/producer-service/internal/usecase"
	"net/http"
)

func main() {

	kafkaConsumer, err := kafka.NewKafkaProducer("localhost:9092")
	if err != nil {
		panic(err)
	}
	exampleUsecase := usecase.NewExampleProducerUseCase(kafkaConsumer)

	producerHandler := webserver.NewProducerHandler(exampleUsecase)

	mux := http.NewServeMux()
	mux.HandleFunc("/", producerHandler.Handle)
	if err := http.ListenAndServe(":3333", mux); err != nil {
		panic(err)
	}
}
