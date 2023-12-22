package statusbar

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
)

type Model struct {
	Status  string
	Mode    string
	Command textinput.Model
}

func New() Model {
	m := Model{
		Status:  "NORMAL",
		Mode:    "Deployments",
		Command: textinput.New(),
	}
	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case ":":
			m.Command.Focus()
		}
	}
	m.Command, cmd = m.Command.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	status := lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("200")).
		Foreground(lipgloss.Color("#333333")).
		Width(8).
		Bold(true).
		Render(m.Status)
	mode := lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#444444")).Render(m.Mode)
	command := lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("#333333")).
		Width(tui.WindowSize.Width - 8).Render(m.Command.View())
	return lipgloss.JoinHorizontal(0, status, mode, command)
}
