package deployments

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
	"os"
	"strings"
)

type FocusMsg bool

type Model struct {
	Ready       bool
	Table       table.Model
	namespaces  []string
	deployments []deployment
}

type deployment struct {
	name      string
	namespace string
}

func New() Model {
	namespaces := getNamespaces()
	deployments := getDeployments(namespaces)

	columns := []table.Column{
		{Title: "Deployment", Width: 32},
		{Title: "Namespace", Width: 20},
	}

	var rows []table.Row
	for _, deployment := range deployments {
		row := table.Row{deployment.name, deployment.namespace}
		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
	)

	ts := table.DefaultStyles()

	ts.Header = ts.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true)

	ts.Selected = ts.Selected.
		Foreground(lipgloss.Color("#FDFDFD")).
		Background(lipgloss.Color("#026873"))

	t.SetStyles(ts)

	m := Model{
		Ready:       false,
		Table:       t,
		deployments: deployments,
		namespaces:  namespaces,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
				return m, func() tea.Msg {
					return FocusMsg(false)
				}
			} else {
				m.Table.Focus()
				return m, func() tea.Msg {
					return FocusMsg(true)
				}
			}
		}
	}

	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Width(tui.WindowSize.Width - 2).
		Height(tui.WindowSize.Height - 9)

	var b strings.Builder

	b.WriteString(m.Table.View())

	return style.Render(b.String())
}

func getNamespaces() []string {
	//TODO: Change how we handle finding the dir
	appsPath := os.Getenv("BOSUN_FLUX_DIR") + "/kubernetes/apps"
	var namespaces []string
	// TODO: Handle this error
	nsDirs, _ := os.ReadDir(appsPath)
	for _, match := range nsDirs {
		if match.IsDir() {
			namespaces = append(namespaces, match.Name())
		}
	}
	return namespaces
}

func getDeployments(namespaces []string) []deployment {
	appsDir := os.Getenv("BOSUN_FLUX_DIR") + "/kubernetes/apps"
	var deployments []deployment
	for _, namespace := range namespaces {
		nsDir, _ := os.ReadDir(appsDir + "/" + namespace)
		for _, match := range nsDir {
			if match.IsDir() {
				deployment := deployment{
					name:      match.Name(),
					namespace: namespace,
				}
				deployments = append(deployments, deployment)
			}
		}
	}
	return deployments
}

// var style = tui.BorderStyle.
//
//	Width(tui.WindowSize.Width - 2).
//	Height(tui.WindowSize.Height - tui.HeaderHeight - 2)
//
// type FocusedMsg bool
//
//	type Model struct {
//		ready       bool
//		width       int
//		height      int
//		Init        bool
//		Table       table.Model
//		namespaces  []string
//		deployments []deployment
//	}
//
//	type deployment struct {
//		name      string
//		namespace string
//	}
//
//	func New() Model {
//		namespaces := getNamespaces()
//		deployments := getDeployments(namespaces)
//
//		columns := []table.Column{
//			{Title: "Deployment", Width: 32},
//			{Title: "Namespace", Width: 20},
//		}
//
//		var rows []table.Row
//		for _, deployment := range deployments {
//			row := table.Row{deployment.name, deployment.namespace}
//			rows = append(rows, row)
//		}
//
//		t := table.New(
//			table.WithColumns(columns),
//			table.WithRows(rows),
//		)
//
//		ts := table.DefaultStyles()
//
//		ts.Header = ts.Header.
//			BorderStyle(lipgloss.NormalBorder()).
//			BorderForeground(lipgloss.Color("240")).
//			BorderBottom(true)
//
//		ts.Selected = ts.Selected.
//			Foreground(lipgloss.Color("#FDFDFD")).
//			Background(lipgloss.Color("#026873"))
//
//		t.SetStyles(ts)
//
//		m := Model{
//			ready:       false,
//			namespaces:  namespaces,
//			deployments: deployments,
//			Table:       t,
//			Init:        false,
//		}
//
//		return m
//	}
//
//	func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
//		// Minus 2 for margins
//		style = tui.BorderStyle.
//			Width(tui.WindowSize.Width - 2).
//			Height(tui.WindowSize.Height - tui.HeaderHeight - 2)
//		var cmd tea.Cmd
//		var cmds []tea.Cmd
//		switch msg := msg.(type) {
//		case tea.KeyMsg:
//			switch msg.String() {
//			case "esc":
//				if m.Table.Focused() {
//					m.Table.Blur()
//					style = style.BorderForeground(lipgloss.Color("212"))
//				} else {
//					m.Table.Focus()
//					style = style.BorderForeground(lipgloss.Color("39"))
//					cmds = append(cmds, func() tea.Msg {
//						return FocusedMsg(true)
//					})
//				}
//			}
//		}
//		m.Table, cmd = m.Table.Update(msg)
//		cmds = append(cmds, cmd)
//		return m, tea.Batch(cmds...)
//	}
//
//	func (m Model) View() string {
//		var b strings.Builder
//		b.WriteString(m.Table.View())
//		return style.Render(b.String())
//	}
