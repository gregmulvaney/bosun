package browser

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
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
	deploymentView state = iota
	generateView
)

type model struct {
	state       state
	ready       bool
	focused     bool
	statusbar   statusbar.Model
	deployments deployments.Model
	generate    generate.Model
}

func Browser() *tea.Program {

	d := deployments.New()
	s := statusbar.New()

	m := model{
		state:       deploymentView,
		deployments: d,
		statusbar:   s,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	return p
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	switch m.state {
	default:

	}
	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case deployments.FocusMsg:
		if msg {
			m.focused = true
		} else {
			m.focused = false
		}
	case generate.FocusMsg:
		if msg {
			m.focused = true
		} else {
			m.focused = false
		}
	case tea.WindowSizeMsg:
		tui.WindowSize = msg

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, tui.Keymap.Quit):
			return m, tea.Quit
		}
		// String key events
		switch msg.String() {
		case "g":
			// Switch to generate view unless already on generate
			if !m.focused && m.state != generateView {
				m.generate = generate.New()
				m.state = generateView
				m.generate.AppData.Inputs[0].Focus()
				m.focused = true
			}
		case "q":
			if !m.focused {
				m.state = deploymentView
			}
		}
	}
	m.statusbar, cmd = m.statusbar.Update(msg)
	cmds = append(cmds, cmd)

	switch m.state {
	case generateView:
		m.generate, cmd = m.generate.Update(msg)
		return m, cmd
	default:
		if !m.deployments.Ready {
			m.deployments.Table.Focus()
			m.focused = true
			m.deployments.Ready = true
		}
		m.deployments, cmd = m.deployments.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(logo)
	b.WriteString("\n\n")

	switch m.state {
	case generateView:
		b.WriteString(m.generate.View())
	case deploymentView:
		b.WriteString(m.deployments.View())
	}
	b.WriteString("\n")
	b.WriteString(m.statusbar.View())
	return b.String()
}
