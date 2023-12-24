package browser

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
	"github.com/gregmulvaney/bosun/internal/tui/components/statusbar"
	"github.com/gregmulvaney/bosun/internal/tui/deployments"
	"github.com/gregmulvaney/bosun/internal/tui/generate"
	"strings"
)

var logo = `   ___  ____  ______  ___  __
  / _ )/ __ \/ __/ / / / |/ /
 / _  / /_/ /\ \/ /_/ /    / 
/____/\____/___/\____/_/|_/`

type state int

const (
	deploymentsView state = iota
	generateView
)

type Model struct {
	state       state
	statusbar   statusbar.Model
	deployments deployments.Model
	generate    generate.Model
}

func Browser() *tea.Program {
	s := statusbar.New()
	d := deployments.New()

	m := Model{
		state:       deploymentsView,
		statusbar:   s,
		deployments: d,
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	return p
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.statusbar.Init())
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		tui.WindowSize = msg
	case statusbar.ModeInsertMsg:
		switch m.state {
		case generateView:
			m.generate.Focus()
		default:
			m.deployments.Table.Focus()
		}
	case statusbar.ModeNormalMsg:
		switch m.state {
		default:
			m.deployments.Table.Blur()
		}
	case statusbar.SpawnGenerateMsg:
		m.generate = generate.New()
		m.state = generateView
	case statusbar.SpawnDeploymentsMsg:
		m.state = deploymentsView
	case generate.ChartGeneratedMsg:
		m.state = deploymentsView
		m.statusbar.Mode = statusbar.Normal
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, tui.Keymap.Quit):
			return m, tea.Quit
		}
	}
	if m.statusbar.Mode == statusbar.Insert {
		switch m.state {
		case generateView:
			m.generate, cmd = m.generate.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.deployments, cmd = m.deployments.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	m.statusbar, cmd = m.statusbar.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder
	logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("226"))
	b.WriteString(logoStyle.Render(logo) + "\n\n")
	switch m.state {
	case generateView:
		b.WriteString(m.generate.View())
	default:
		b.WriteString(m.deployments.View())
	}
	b.WriteString("\n")
	b.WriteString(m.statusbar.View())
	return b.String()
}
