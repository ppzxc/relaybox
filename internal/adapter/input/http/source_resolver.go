package http

import "relaybox/internal/domain"

// SourceResolver URL sourceID를 domain.SourceType으로 변환하고 토큰을 검증한다.
type SourceResolver interface {
	Resolve(sourceID string) (domain.SourceType, error)
	ValidateToken(sourceID, token string) bool
}
