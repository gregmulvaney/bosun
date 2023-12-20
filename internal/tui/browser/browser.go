package browser

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gregmulvaney/bosun/internal/tui"
	"github.com/gregmulvaney/bosun/internal/tui/deployments"
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

type model struct {
	state       state
	ready       bool
	width       int
	viewport    viewport.Model
	deployments deployments.Model
}

func Browser() *tea.Program {
	d := deployments.New()

	m := model{
		state:       deploymentsView,
		ready:       false,
		deployments: d,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	return p
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	//var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		tui.WindowSize = msg
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-tui.HeaderHeight)
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - tui.HeaderHeight
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, tui.Keymap.Quit):
			return m, tea.Quit
		}
		switch msg.String() {
		case "g":
			m.state = generateView
		case "q":
			m.state = deploymentsView
		}

	}
	switch m.state {
	default:
		m.deployments.Table.Focus()
		m.deployments, cmd = m.deployments.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString(logo)
	b.WriteString("\n\n")
	switch m.state {
	case generateView:
		m.viewport.SetContent("Generate")
		b.WriteString(tui.BorderStyle.Width(tui.WindowSize.Width - 2).Render(m.viewport.View()))
		return b.String()
	default:
		m.viewport.SetContent(m.deployments.View())
		b.WriteString(m.viewport.View())
		return b.String()
	}
}

//const (
//	deploymentsView state = iota
//	generate
//)
//
//type model struct {
//	ready           bool
//	width           int
//	height          int
//	state           state
//	contentViewport viewport.Model
//	deployments     deployments.Model
//}
//
//func Browser() *tea.Program {
//	d := deployments.New()
//	m := model{
//		ready:       false,
//		deployments: d,
//		state:       deploymentsView,
//	}
//
//	p := tea.NewProgram(m, tea.WithAltScreen())
//	return p
//}
//
//func (m model) Init() tea.Cmd {
//	return nil
//}
//
//func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch msg.String() {
//		case "ctrl+c":
//			return m, tea.Quit
//		}
//	case tea.WindowSizeMsg:
//		if !m.ready {
//			m.contentViewport = viewport.New(msg.Width, msg.Height-7)
//			m.width = msg.Width
//		} else {
//			m.contentViewport.Height = msg.Height - 7
//			m.contentViewport.Width = msg.Width
//			m.width = msg.Width
//		}
//
//	}
//	return m, nil
//}
//
//func (m model) View() string {
//	var b strings.Builder
//	b.WriteString(PrimaryColor.Padding(1, 1).Render(logo) + "\n")
//	switch m.state {
//	default:
//		m.contentViewport.SetContent(m.deployments.View())
//	}
//	b.WriteString(lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#F27405")).Render(m.contentViewport.View()))
//	return b.String()
//}
