package tui

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor  = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	ViewportStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("149")).Padding(0, 2)
)
