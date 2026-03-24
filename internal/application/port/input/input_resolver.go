package input

// InputResolver URL inputID를 검증하고 토큰을 확인한다.
type InputResolver interface {
	Resolve(inputID string) (string, error)
	ValidateToken(inputID, token string) bool
}
