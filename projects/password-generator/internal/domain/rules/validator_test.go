package rules

import "testing"

func TestValidatePasswordConfig_ValidConfig(t *testing.T) {
	cfg := PasswordConfig{Length: 16, Uppercase: true, Lowercase: true, Digits: true, Symbols: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_InvalidLength(t *testing.T) {
	cfg := PasswordConfig{Length: 4, Lowercase: true}

	if err := ValidatePasswordConfig(cfg); err == nil {
		t.Fatalf("expected length validation error")
	}
}

func TestValidatePasswordConfig_NoCategorySelected(t *testing.T) {
	cfg := PasswordConfig{Length: 12}

	if err := ValidatePasswordConfig(cfg); err == nil {
		t.Fatalf("expected category validation error")
	}
}
