package output

import (
	"context"

	"relaybox/internal/domain"
)

type RouteConfigReader interface {
	GetChannels(ctx context.Context, sourceID string) ([]domain.Channel, error)
}
