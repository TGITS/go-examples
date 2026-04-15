package password

import (
	"testing"

	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/domain/rules"
)

func TestEvaluateStrength_Weak(t *testing.T) {
	cfg := rules.PasswordConfig{Length: 8, Lowercase: true}
	if got := EvaluateStrength("abcdefgh", cfg); got != "Weak" {
		t.Fatalf("expected Weak, got %s", got)
	}
}

func TestEvaluateStrength_Strong(t *testing.T) {
	cfg := rules.PasswordConfig{Length: 16, Uppercase: true, Lowercase: true, Digits: true, Symbols: true}
	if got := EvaluateStrength("Ab1!Ab1!Ab1!Ab1!", cfg); got != "Strong" {
		t.Fatalf("expected Strong, got %s", got)
	}
}
