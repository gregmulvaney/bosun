package generate

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
	"github.com/gregmulvaney/bosun/internal/tui/components/textform"
)

type FocusMsg bool

type state int

const (
	appDataView state = iota
)

type Model struct {
	state        state
	formViewport viewport.Model
	AppData      textform.Model
}

func New() Model {
	appDataFields := []string{"App name", "Namespace", "Repository", "Tag"}
	appDataForm := textform.New(appDataFields, "Enter app details")

	formView := viewport.New(tui.WindowSize.Width/2, tui.WindowSize.Height-9)
	m := Model{
		state:        appDataView,
		formViewport: formView,
		AppData:      appDataForm,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.formViewport.Width = tui.WindowSize.Width / 2
		m.formViewport.Height = tui.WindowSize.Height - 9
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			switch m.state {
			default:
				if m.AppData.Inputs[m.AppData.FocusIndex].Focused() {
					m.AppData.Inputs[m.AppData.FocusIndex].Blur()
					return m, func() tea.Msg {
						return FocusMsg(false)
					}
				} else {
					m.AppData.Inputs[m.AppData.FocusIndex].Focus()
					return m, func() tea.Msg {
						return FocusMsg(true)
					}
				}
			}
		}
	}

	var cmd tea.Cmd
	switch m.state {
	default:
		m.AppData, cmd = m.AppData.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39"))

	switch m.state {
	default:
		m.formViewport.SetContent(m.AppData.View())
		return style.Padding(1, 2).Width(tui.WindowSize.Width / 2).Render(m.formViewport.View())
	}
}

//type state int
//
//const (
//	appDataView state = iota
//)
//
//type Model struct {
//	state        state
//	ready        bool
//	height       int
//	width        int
//	formViewport viewport.Model
//	codeViewport viewport.Model
//	AppData      textform.Model
//}
//
//func New() Model {
//	appDataFields := []string{"App name", "Namespace", "Repository", "Tag"}
//	appDataForm := textform.New(appDataFields, "Enter your apps details")
//
//	formViewport := viewport.New((tui.WindowSize.Width/3)-2, tui.WindowSize.Height-tui.HeaderHeight-3)
//	formViewport.SetContent(appDataForm.View())
//
//	m := Model{
//		state:   appDataView,
//		ready:   false,
//		AppData: appDataForm,
//	}
//
//	return m
//}
//
//func (m Model) Init() tea.Cmd {
//	var cmds []tea.Cmd
//	switch m.state {
//	case appDataView:
//		cmds = append(cmds, m.AppData.Init())
//		return tea.Batch(cmds...)
//	}
//	return nil
//}
//
//func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch msg.String() {
//		case "ctrl+c":
//			return m, tea.Quit
//		}
//
//	case tea.WindowSizeMsg:
//		if !m.ready {
//			m.ready = true
//		} else {
//			m.formViewport.Width = (tui.WindowSize.Width / 3) - 4
//			m.formViewport.Height = tui.WindowSize.Height - tui.HeaderHeight - 3
//		}
//	}
//	var cmd tea.Cmd
//	switch m.state {
//	case appDataView:
//		m.AppData, cmd = m.AppData.Update(msg)
//		return m, cmd
//	}
//	return m, nil
//}
//
//func (m Model) View() string {
//	switch m.state {
//	default:
//		m.formViewport.SetContent(m.AppData.View())
//	}
//	return tui.BorderStyle.Width(tui.WindowSize.Width / 3).Render(m.formViewport.View())
//}
//
//func (m Model) renderView(formViewport viewport.Model, codeViewport viewport.Model) string {
//	return lipgloss.Place(m.width, m.height, 0, 0, lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.NewStyle().Render(formViewport.View()), lipgloss.NewStyle().Render(codeViewport.View())))
//}
