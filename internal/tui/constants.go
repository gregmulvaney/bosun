package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	WindowSize tea.WindowSizeMsg
)

var PrimaryColor = lipgloss.Color("39")
var SecondaryColor = lipgloss.Color("40")

var HeaderHeight = 5

var BorderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(PrimaryColor)

type keymap struct {
	Quit key.Binding
}

var Keymap = keymap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "Quit"),
	),
}
