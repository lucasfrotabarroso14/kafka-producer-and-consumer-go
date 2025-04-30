package usecase

import (
	"context"
	"fmt"
)

type ProcessMessageUseCase struct {
}

func NewProcessMessageUseCase() *ProcessMessageUseCase {
	return &ProcessMessageUseCase{}
}

func (p *ProcessMessageUseCase) Execute(ctx context.Context, key, value []byte) error {

	fmt.Println("usecase - processando mensagem: %s => %s \n", string(key), string(value))
	return nil
}
