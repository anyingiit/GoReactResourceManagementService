package globalVars

import (
	"testing"
)

func TestGlobalVars_Set(t *testing.T) {
	g := newGlobalVars("test")
	err := g.Set(42)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	err = g.Set(43)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}

func TestGlobalVars_Get(t *testing.T) {
	g := newGlobalVars("test")
	val, err := g.Get()
	if val != nil {
		t.Errorf("expected nil value, but got %v", val)
	}
	if err == nil {
		t.Error("expected error, but got nil")
	}
	g.Set(42)
	val, err = g.Get()
	if val != 42 {
		t.Errorf("expected value 42, but got %v", val)
	}
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}
