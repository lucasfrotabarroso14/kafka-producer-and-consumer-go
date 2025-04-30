package port

import "context"

type ProcessTransactionInputDTO struct {
	Message string
}
type ProcessMessageUseCaseInterface interface {
	Execute(ctx context.Context, input ProcessTransactionInputDTO) error
}
