package logging

import (
	"context"
	"log/slog"
)

// MultiHandler는 여러 slog.Handler에 로그 레코드를 fanout한다.
type MultiHandler struct {
	handlers []slog.Handler
}

// NewMultiHandler는 주어진 handler들을 묶는 MultiHandler를 반환한다.
func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

// Enabled는 하나 이상의 handler가 해당 레벨을 처리할 수 있으면 true를 반환한다.
func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle은 활성화된 각 handler에 레코드 복사본을 전달한다.
// r.Clone()을 사용해 data race를 방지한다.
func (m *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range m.handlers {
		if h.Enabled(ctx, r.Level) {
			if err := h.Handle(ctx, r.Clone()); err != nil {
				return err
			}
		}
	}
	return nil
}

// WithAttrs는 각 handler에 attrs를 전파한 새 MultiHandler를 반환한다.
func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	hs := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		hs[i] = h.WithAttrs(attrs)
	}
	return NewMultiHandler(hs...)
}

// WithGroup는 각 handler에 group을 전파한 새 MultiHandler를 반환한다.
func (m *MultiHandler) WithGroup(name string) slog.Handler {
	hs := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		hs[i] = h.WithGroup(name)
	}
	return NewMultiHandler(hs...)
}
