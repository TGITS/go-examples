package rules

import "testing"

func TestValidatePasswordConfig_ValidConfig(t *testing.T) {
	cfg := PasswordConfig{Length: 16, Uppercase: true, Lowercase: true, Digits: true, Symbols: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_InvalidMinLength(t *testing.T) {
	cfg := PasswordConfig{Length: 7, Lowercase: true}

	if err := ValidatePasswordConfig(cfg); err == nil {
		t.Fatalf("expected length validation error")
	}
}

func TestValidatePasswordConfig_InvalidMaxLength(t *testing.T) {
	cfg := PasswordConfig{Length: 129, Symbols: true}

	if err := ValidatePasswordConfig(cfg); err == nil {
		t.Fatalf("expected length validation error")
	}
}

func TestValidatePasswordConfig_ZeroLength(t *testing.T) {
	cfg := PasswordConfig{Length: 0, Lowercase: true}

	if err := ValidatePasswordConfig(cfg); err == nil {
		t.Fatalf("expected length validation error for zero length")
	}
}

func TestValidatePasswordConfig_NegativeLength(t *testing.T) {
	cfg := PasswordConfig{Length: -1, Lowercase: true}

	if err := ValidatePasswordConfig(cfg); err == nil {
		t.Fatalf("expected length validation error for negative length")
	}
}

func TestValidatePasswordConfig_NoCategorySelected(t *testing.T) {
	cfg := PasswordConfig{Length: 12}

	if err := ValidatePasswordConfig(cfg); err == nil {
		t.Fatalf("expected category validation error")
	}
}

func TestValidatePasswordConfig_LengthAtMinIncluded(t *testing.T) {
	cfg := PasswordConfig{Length: 8, Uppercase: true, Lowercase: true, Digits: true, Symbols: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_LengthAtMaxIncluded(t *testing.T) {
	cfg := PasswordConfig{Length: 128, Uppercase: true, Lowercase: true, Digits: true, Symbols: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_LengthNearBoundariesIncluded(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{name: "min plus one", length: MinLength + 1},
		{name: "max minus one", length: MaxLength - 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := PasswordConfig{Length: tt.length, Lowercase: true}

			if err := ValidatePasswordConfig(cfg); err != nil {
				t.Fatalf("expected no error for length %d, got %v", tt.length, err)
			}
		})
	}
}

func TestValidatePasswordConfig_OnlyUpperCaseEnough(t *testing.T) {
	cfg := PasswordConfig{Length: 12, Uppercase: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_OnlyLowerCaseEnough(t *testing.T) {
	cfg := PasswordConfig{Length: 12, Lowercase: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_OnlyDigitsEnough(t *testing.T) {
	cfg := PasswordConfig{Length: 12, Digits: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_OnlySymbolsEnough(t *testing.T) {
	cfg := PasswordConfig{Length: 12, Symbols: true}

	if err := ValidatePasswordConfig(cfg); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidatePasswordConfig_InvalidMinLength_ErrorMessageForInvalidMin(t *testing.T) {
	cfg := PasswordConfig{Length: 4, Lowercase: true}
	err := ValidatePasswordConfig(cfg)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	want := "length must be between 8 and 128"
	if err.Error() != want {
		t.Fatalf("expected error message %q, got %q", want, err.Error())
	}
}

func TestValidatePasswordConfig_InvalidMaxLength_ErrorMessageForInvalidMax(t *testing.T) {
	cfg := PasswordConfig{Length: 129, Symbols: true}
	err := ValidatePasswordConfig(cfg)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	want := "length must be between 8 and 128"
	if err.Error() != want {
		t.Fatalf("expected error message %q, got %q", want, err.Error())
	}
}

func TestValidatePasswordConfig_NoCategorySelected_ErrorMessageForNoCategory(t *testing.T) {
	cfg := PasswordConfig{Length: 12}
	err := ValidatePasswordConfig(cfg)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	want := "at least one character category must be enabled"
	if err.Error() != want {
		t.Fatalf("expected error message %q, got %q", want, err.Error())
	}
}

func TestValidatePasswordConfig_DoubleInvalidPrioritizesLengthError(t *testing.T) {
	cfg := PasswordConfig{Length: 0}
	err := ValidatePasswordConfig(cfg)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	validationErrs, ok := err.(ValidationErrors)
	if !ok {
		t.Fatalf("expected ValidationErrors, got %T", err)
	}

	if len(validationErrs.Messages) != 2 {
		t.Fatalf("expected 2 validation messages, got %d", len(validationErrs.Messages))
	}

	wantMessages := []string{
		"length must be between 8 and 128",
		"at least one character category must be enabled",
	}

	for index, want := range wantMessages {
		if validationErrs.Messages[index] != want {
			t.Fatalf("expected validation message %d to be %q, got %q", index, want, validationErrs.Messages[index])
		}
	}
}
