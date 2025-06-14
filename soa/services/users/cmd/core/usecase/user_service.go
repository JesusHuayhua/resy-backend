package usecase

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, userId string)
	Status(ctx context.Context, userId string)
	ServiceStatus(ctx context.Context) (int, error)
}
