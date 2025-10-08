package health

import (
	"context"
)

type Repository interface {
	HealthCheck(ctx context.Context) error
}
