package password

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/domain/rules"
)

const (
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	digitChars     = "0123456789"
	symbolChars    = "!@#$%^&*()-_=+[]{}<>?/|"
)

func Generate(cfg rules.PasswordConfig) (string, error) {
	if err := rules.ValidatePasswordConfig(cfg); err != nil {
		return "", err
	}

	selectedSets := buildSelectedSets(cfg)
	if len(selectedSets) == 0 {
		return "", fmt.Errorf("no character set selected")
	}

	combined := strings.Join(selectedSets, "")
	if combined == "" {
		return "", fmt.Errorf("combined character set is empty")
	}

	result := make([]byte, 0, cfg.Length)

	for _, set := range selectedSets {
		c, err := randomChar(set)
		if err != nil {
			return "", err
		}
		result = append(result, c)
	}

	for len(result) < cfg.Length {
		c, err := randomChar(combined)
		if err != nil {
			return "", err
		}
		result = append(result, c)
	}

	if err := secureShuffle(result); err != nil {
		return "", err
	}

	return string(result), nil
}

func buildSelectedSets(cfg rules.PasswordConfig) []string {
	sets := make([]string, 0, 4)

	if cfg.Uppercase {
		sets = append(sets, uppercaseChars)
	}
	if cfg.Lowercase {
		sets = append(sets, lowercaseChars)
	}
	if cfg.Digits {
		sets = append(sets, digitChars)
	}
	if cfg.Symbols {
		sets = append(sets, symbolChars)
	}

	return sets
}

func randomChar(set string) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
	if err != nil {
		return 0, fmt.Errorf("secure random failed: %w", err)
	}

	return set[n.Int64()], nil
}

func secureShuffle(data []byte) error {
	for i := len(data) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return fmt.Errorf("secure shuffle failed: %w", err)
		}

		j := int(n.Int64())
		data[i], data[j] = data[j], data[i]
	}

	return nil
}
