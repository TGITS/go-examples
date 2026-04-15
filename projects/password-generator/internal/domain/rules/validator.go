package rules

import (
	"fmt"
	"strings"
)

const (
	MinLength = 8
	MaxLength = 128
)

type PasswordConfig struct {
	Length        int
	Uppercase     bool
	Lowercase     bool
	Digits        bool
	Symbols       bool
	Exclude       string
	BatchSize     int
	AutoCopy      bool
	SaveToFile    bool
	OutputFile    string
	UseEntropy    bool
	UsePassphrase bool
}

type ValidationErrors struct {
	Messages []string
}

func (e ValidationErrors) Error() string {
	return strings.Join(e.Messages, "; ")
}

func ValidatePasswordConfig(cfg PasswordConfig) error {
	messages := make([]string, 0, 2)

	if cfg.Length < MinLength || cfg.Length > MaxLength {
		messages = append(messages, fmt.Sprintf("length must be between %d and %d", MinLength, MaxLength))
	}

	if !cfg.Uppercase && !cfg.Lowercase && !cfg.Digits && !cfg.Symbols {
		messages = append(messages, "at least one character category must be enabled")
	}

	if len(messages) > 0 {
		return ValidationErrors{Messages: messages}
	}

	return nil
}
