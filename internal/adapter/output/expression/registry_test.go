package expression_test

import (
	"testing"

	"relaybox/internal/adapter/output/expression"
	"relaybox/internal/application/port/output"
)

var _ output.ExpressionEngineRegistry = (*expression.InMemoryExpressionEngineRegistry)(nil)

func TestRegistry_GetAndDefault(t *testing.T) {
	reg := expression.NewInMemoryExpressionEngineRegistry()
	celEng := newCELEngine(t)
	exprEng := expression.NewExprEngine()

	reg.Register(celEng)
	reg.Register(exprEng)

	// Default is first registered
	if reg.Default().Type() != "CEL" {
		t.Errorf("Default() = %q, want CEL", reg.Default().Type())
	}

	got, err := reg.Get("EXPR")
	if err != nil {
		t.Fatalf("Get(EXPR) error: %v", err)
	}
	if got.Type() != "EXPR" {
		t.Errorf("Get(EXPR).Type() = %q", got.Type())
	}

	_, err = reg.Get("unknown")
	if err == nil {
		t.Error("expected error for unknown engine")
	}
}

func TestRegistry_SetDefault(t *testing.T) {
	reg := expression.NewInMemoryExpressionEngineRegistry()
	reg.Register(newCELEngine(t))
	reg.Register(expression.NewExprEngine())

	if err := reg.SetDefault("EXPR"); err != nil {
		t.Fatalf("SetDefault error: %v", err)
	}
	if reg.Default().Type() != "EXPR" {
		t.Errorf("Default() = %q, want EXPR", reg.Default().Type())
	}

	if err := reg.SetDefault("nonexistent"); err == nil {
		t.Error("expected error for nonexistent engine")
	}
}

func TestRegistry_EmptyDefault(t *testing.T) {
	reg := expression.NewInMemoryExpressionEngineRegistry()
	if reg.Default() != nil {
		t.Error("Default() on empty registry should be nil")
	}
}
