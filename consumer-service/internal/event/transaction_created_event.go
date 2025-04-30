package event

import "async_kafka_services/pkg/events"

func NewTransactionCreatedEvent(message string) *events.BaseEvent {
	return events.NewBaseEvent("TransactionCreated", message)
}
