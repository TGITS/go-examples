package rules

import "fmt"

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

func ValidatePasswordConfig(cfg PasswordConfig) error {
	if cfg.Length < MinLength || cfg.Length > MaxLength {
		return fmt.Errorf("length must be between %d and %d", MinLength, MaxLength)
	}

	if !cfg.Uppercase && !cfg.Lowercase && !cfg.Digits && !cfg.Symbols {
		return fmt.Errorf("at least one character category must be enabled")
	}

	return nil
}
