package app

import "fmt"

func Render(m Model) string {
	if m.LastError != "" {
		return fmt.Sprintf(
			"Password Generator\n\nError: %s\n\nPress %s or %s to generate, %s to quit.\n",
			m.LastError,
			KeyGenerateA,
			KeyGenerateB,
			KeyQuitA,
		)
	}

	if m.Password == "" {
		return fmt.Sprintf(
			"Password Generator\n\nNo password yet.\n\nPress %s or %s to generate, %s to quit.\n",
			KeyGenerateA,
			KeyGenerateB,
			KeyQuitA,
		)
	}

	return fmt.Sprintf(
		"Password Generator\n\nPassword: %s\nStrength: %s\n\nPress %s or %s to generate again, %s to quit.\n",
		m.Password,
		m.Strength,
		KeyGenerateA,
		KeyGenerateB,
		KeyQuitA,
	)
}
