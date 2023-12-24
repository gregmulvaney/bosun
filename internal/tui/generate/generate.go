package generate

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
	"github.com/gregmulvaney/bosun/internal/tui/components/textform"
	"strings"
)

type ChartGeneratedMsg bool

type state int

const (
	appDataView state = iota
)

type Model struct {
	state     state
	chart     HelmChart
	appData   textform.Model
	namespace string
}

func New() Model {
	appDataFields := []string{"AppName", "Namespace", "Repository", "Tag"}
	appDataForm := textform.New(appDataFields, "Enter your app details")

	chart := NewChart()

	m := Model{
		state:   appDataView,
		chart:   chart,
		appData: appDataForm,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg.(type) {
	case textform.CompleteMsg:
		switch m.state {
		case appDataView:
			values := m.appData.Values()
			m.chart.Metadata.Name = values["AppName"]
			m.namespace = values["Namespace"]
			m.chart.Spec.Values.Controllers.Main.Containers.Main.Image = ContainerImage{
				Repository: values["Repository"],
				Tag:        values["Tag"],
			}
			Build(m.chart, m.namespace)
			return m, func() tea.Msg {
				return ChartGeneratedMsg(true)
			}
		}

	}

	switch m.state {
	default:
		m.appData, cmd = m.appData.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Width(tui.WindowSize.Width-2).
		Height(tui.WindowSize.Height-9).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("39")).
		Padding(1, 1)

	var b strings.Builder

	switch m.state {
	default:
		b.WriteString(m.appData.View())
	}

	return style.Render(b.String())
}

func (m Model) Focus() {
	m.appData.Inputs[m.appData.FocusIndex].Focus()
}
