package config

import "github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/domain/rules"

func DefaultPasswordConfig() rules.PasswordConfig {
	return rules.PasswordConfig{
		Length:    12,
		Uppercase: true,
		Lowercase: true,
		Digits:    true,
		Symbols:   true,
	}
}
