package main

import (
	event2 "async_kafka_services/consumer-service/internal/event"
	"async_kafka_services/consumer-service/internal/event/handler"
	"fmt"
	"os"
	"os/signal"

	"async_kafka_services/consumer-service/internal/infra/kafka"
	"async_kafka_services/consumer-service/internal/usecase"
	"async_kafka_services/pkg/events"
)

func main() {

	consumer, err := kafka.NewKafkaConsumer("localhost:9092", "client-consumer-group")
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	exampleUsecase := usecase.NewProcessTransactionUseCase()
	transactionHandler := handler.NewTransactionCreatedHandler(exampleUsecase)

	dispatcher := events.NewEventDispatcher()
	if err = dispatcher.Register("TransactionCreated", transactionHandler); err != nil {
		panic(err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		err = consumer.Consume("transactions", func(key []byte, value []byte) {
			payload := string(value)
			event := event2.NewTransactionCreatedEvent(payload)
			if err = dispatcher.Dispatch(event); err != nil {
				fmt.Println("erro ao despachar evento:", err)
			}
		})
		if err != nil {
			fmt.Println("Erro no consumer:", err)
		}
	}()
	fmt.Println("Consumer rodando. Pressione Ctrl+C para encerrar...")

	if err != nil {
		panic(err)
	}

	<-stopChan
	fmt.Println("Encerrando consumer.")
	_ = consumer.Close()

}
