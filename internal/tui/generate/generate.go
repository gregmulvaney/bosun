package generate

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
}

func New() Model {
	m := Model{}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

	return m, nil
}

func (m Model) View() string {
	return "Generate"
}
