package usecase

import (
	"context"
	"soa/services/users/pkg/core/internal"
)

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter)
	Status(ctx context.Context, userId string)
	ServiceStatus(ctx context.Context) (int, error)
}
