package port

import (
	"context"
)

type ExampleUseCase interface {
	Execute(ctx context.Context, input InputDTO) (int, error)
}
