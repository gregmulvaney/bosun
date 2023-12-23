package deployments

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
	"os"
)

type Model struct {
	Table       table.Model
	namespaces  []string
	deployments []deployment
}

type deployment struct {
	name      string
	namespace string
}

func New() Model {
	ns := getNamespaces()
	deployments := getDeployments(ns)
	cols := []table.Column{
		{Title: "Deployment", Width: 32},
		{Title: "Namespace", Width: 20},
	}
	var rows []table.Row
	for _, deployment := range deployments {
		rows = append(rows, table.Row{deployment.name, deployment.namespace})
	}

	t := table.New(table.WithColumns(cols), table.WithRows(rows))

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
		Table:       t,
		namespaces:  ns,
		deployments: deployments,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Width(tui.WindowSize.Width - 2).
		Height(tui.WindowSize.Height - 9)

	return style.Render(m.Table.View())
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
