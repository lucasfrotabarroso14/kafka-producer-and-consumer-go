package handler

import (
	"async_kafka_services/consumer-service/internal/domain/port"
	"async_kafka_services/pkg/events"
	"context"
	"fmt"
	"sync"
)

type TransactionCreatedHandler struct {
	ProcessUsecase port.ProcessMessageUseCaseInterface
}

func NewTransactionCreatedHandler(processUsecase port.ProcessMessageUseCaseInterface) *TransactionCreatedHandler {
	return &TransactionCreatedHandler{
		ProcessUsecase: processUsecase,
	}
}

func (h *TransactionCreatedHandler) Handle(ctx context.Context, event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	payload := event.GetPayload().(string)

	input := port.ProcessTransactionInputDTO{
		payload,
	}

	if err := h.ProcessUsecase.Execute(ctx, input); err != nil {
		fmt.Println("Error processing transaction:", err)
		return
	}
	return

}
