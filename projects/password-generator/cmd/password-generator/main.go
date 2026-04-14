package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/app"
	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/config"
)

func main() {
	model := app.NewModel(config.DefaultPasswordConfig())
	p := tea.NewProgram(model)

	if err := p.Start(); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
