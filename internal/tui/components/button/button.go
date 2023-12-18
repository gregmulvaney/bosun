package button

import "github.com/charmbracelet/lipgloss"

var buttonStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("149"))

type Model struct {
	label      string
	focusIndex int
	position   int
}

func New(label string, focusIndex int, position int) Model {
	m := Model{
		label:      label,
		focusIndex: focusIndex,
		position:   position,
	}

	return m
}

func (m Model) View() string {
	if m.focusIndex == m.position {
		return buttonStyle.Render("[ " + m.label + " ]")
	} else {
		return "[ " + m.label + " ]"
	}
}
