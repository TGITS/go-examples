package password

import (
	"strings"
	"testing"

	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/domain/rules"
)

func TestGenerate_ReturnsExpectedLength(t *testing.T) {
	cfg := rules.PasswordConfig{Length: 20, Uppercase: true, Lowercase: true, Digits: true, Symbols: true}

	pwd, err := Generate(cfg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(pwd) != 20 {
		t.Fatalf("expected length 20, got %d", len(pwd))
	}
}

func TestGenerate_ContainsEachEnabledCategory(t *testing.T) {
	cfg := rules.PasswordConfig{Length: 16, Uppercase: true, Lowercase: true, Digits: true, Symbols: true}

	pwd, err := Generate(cfg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assertContainsAny(t, pwd, uppercaseChars)
	assertContainsAny(t, pwd, lowercaseChars)
	assertContainsAny(t, pwd, digitChars)
	assertContainsAny(t, pwd, symbolChars)
}

func TestGenerate_InvalidConfigReturnsError(t *testing.T) {
	cfg := rules.PasswordConfig{Length: 3, Lowercase: true}

	if _, err := Generate(cfg); err == nil {
		t.Fatalf("expected validation error")
	}
}

func assertContainsAny(t *testing.T, value, charset string) {
	t.Helper()

	for i := 0; i < len(charset); i++ {
		if strings.ContainsRune(value, rune(charset[i])) {
			return
		}
	}

	t.Fatalf("expected password to contain at least one character from set %q", charset)
}
