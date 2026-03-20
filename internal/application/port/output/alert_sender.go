package output

import (
	"context"

	"relaybox/internal/domain"
)

type AlertSender interface {
	Send(ctx context.Context, channel domain.Channel, alert domain.Alert) error
}
