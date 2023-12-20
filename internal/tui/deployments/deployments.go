package deployments

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
	"os"
	"strings"
)

type Model struct {
	ready       bool
	width       int
	height      int
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
		ready:       false,
		namespaces:  namespaces,
		deployments: deployments,
		Table:       t,
	}

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		}
	}
	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	style := tui.BorderStyle.Width(tui.WindowSize.Width - 2).Height(tui.WindowSize.Height - tui.HeaderHeight - 2)
	var b strings.Builder
	b.WriteString(m.Table.View())
	return style.Render(b.String())
}

func getNamespaces() []string {
	//TODO: Change how we handle finding the dir
	appsDir := os.Getenv("BOSUN_FLUX_DIR") + "/kubernetes/apps"
	var namespaces []string
	// TODO: Handle this error
	nsDirs, _ := os.ReadDir(appsDir)
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
