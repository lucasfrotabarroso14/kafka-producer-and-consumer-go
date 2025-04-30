package usecase

import (
	"async_kafka_services/consumer-service/internal/domain/port"
	"context"
	"fmt"
)

type ProcessTransactionUseCase struct {
}

func NewProcessTransactionUseCase() *ProcessTransactionUseCase {
	return &ProcessTransactionUseCase{}
}

func (p *ProcessTransactionUseCase) Execute(ctx context.Context, input port.ProcessTransactionInputDTO) error {

	fmt.Println("usecase - processando mensagem: ", input.Message)
	return nil
}
