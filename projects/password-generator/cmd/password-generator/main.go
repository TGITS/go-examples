package main

import (
	"log"

	tea "charm.land/bubbletea/v2"

	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/app"
	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/config"
)

func main() {
	model := app.NewModel(config.DefaultPasswordConfig())
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
