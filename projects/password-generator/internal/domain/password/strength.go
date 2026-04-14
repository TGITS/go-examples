package password

import "github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/domain/rules"

func EvaluateStrength(pwd string, cfg rules.PasswordConfig) string {
	score := 0

	if len(pwd) >= 8 {
		score++
	}
	if len(pwd) >= 12 {
		score++
	}
	if cfg.Uppercase {
		score++
	}
	if cfg.Lowercase {
		score++
	}
	if cfg.Digits {
		score++
	}
	if cfg.Symbols {
		score++
	}

	if score <= 3 {
		return "Weak"
	}
	if score <= 5 {
		return "Medium"
	}

	return "Strong"
}
