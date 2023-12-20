package generate

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui/components/textform"
)

type state int

const (
	appDataView state = iota
)

type model struct {
	state        state
	ready        bool
	height       int
	width        int
	formViewport viewport.Model
	codeViewport viewport.Model
	appData      textform.Model
}

func Prompt(fullscreen bool) *tea.Program {
	appDataFields := []string{"App name", "Namespace", "Repository", "Tag"}
	appDataForm := textform.New(appDataFields, "Enter your apps details")

	m := model{
		state:   appDataView,
		ready:   false,
		appData: appDataForm,
	}

	var p *tea.Program
	if fullscreen {
		p = tea.NewProgram(m, tea.WithAltScreen())
	} else {
		p = tea.NewProgram(m)
	}

	return p
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	switch m.state {
	case appDataView:
		cmds = append(cmds, m.appData.Init())
		return tea.Batch(cmds...)
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		if !m.ready {
			m.formViewport = viewport.New(msg.Width/3, msg.Height-3)
			m.codeViewport = viewport.New(msg.Width/2, msg.Height-3)
			m.height = msg.Height
			m.width = msg.Width
			m.ready = true
		} else {
			m.formViewport.Height = msg.Height - 3
			m.formViewport.Width = msg.Width / 3
			m.codeViewport.Height = msg.Height - 3
			m.codeViewport.Width = msg.Width / 2
			m.width = msg.Width
			m.height = msg.Height
		}
	}
	var cmd tea.Cmd
	switch m.state {
	case appDataView:
		m.appData, cmd = m.appData.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	switch m.state {
	default:
		m.formViewport.SetContent(m.appData.View())
		m.codeViewport.SetContent(lipgloss.NewStyle().Width(40).Render("Code preview placeholder"))
		return m.renderView(m.formViewport, m.codeViewport)
	}
}

func (m model) renderView(formViewport viewport.Model, codeViewport viewport.Model) string {
	return lipgloss.Place(m.width, m.height, 0, 0, lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.NewStyle().Render(formViewport.View()), lipgloss.NewStyle().Render(codeViewport.View())))
}
