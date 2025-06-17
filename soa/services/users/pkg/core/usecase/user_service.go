package usecase

import (
	"context"
	"soa/services/users/pkg/core/svc_internal"
)

type Service interface {
	Get(ctx context.Context, filters ...svc_internal.Filter) error
	Status(ctx context.Context, userId string) (svc_internal.StatusCode, error)
	ServiceStatus(ctx context.Context) (int, error)
}
