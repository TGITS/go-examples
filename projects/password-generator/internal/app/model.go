package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/domain/password"
	"github.com/TGITS/go-examples/go-examples/projects/password-generator/internal/domain/rules"
)

type Model struct {
	Config    rules.PasswordConfig
	Password  string
	Strength  string
	LastError string
}

func NewModel(cfg rules.PasswordConfig) Model {
	return Model{Config: cfg}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case KeyQuitA, KeyQuitB:
			return m, tea.Quit
		case KeyGenerateA, KeyGenerateB:
			pwd, err := password.Generate(m.Config)
			if err != nil {
				m.LastError = err.Error()
				return m, nil
			}
			m.Password = pwd
			m.Strength = password.EvaluateStrength(pwd, m.Config)
			m.LastError = ""
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	return tea.NewView("Password Generator")
}
