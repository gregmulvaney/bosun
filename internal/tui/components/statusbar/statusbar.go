package statusbar

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
)

type ModeInsertMsg bool
type ModeNormalMsg bool
type SpawnGenerateMsg bool
type SpawnDeploymentsMsg bool

type mode int

const (
	Normal mode = iota
	Insert
	Command
)

type Model struct {
	Mode         mode
	CommandInput textinput.Model
}

func New() Model {
	cmdInput := textinput.New()
	cmdInput.Prompt = ""

	m := Model{
		Mode:         Normal,
		CommandInput: cmdInput,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "i":
			if m.Mode == Normal {
				m.Mode = Insert
				return m, func() tea.Msg {
					return ModeInsertMsg(true)
				}
			}
		case ":":
			if !m.CommandInput.Focused() && m.Mode == Normal {
				m.CommandInput.Focus()
				m.Mode = Command
			}
		case "esc":
			if m.Mode == Command || m.Mode == Insert {
				m.Mode = Normal
				m.CommandInput.Blur()
				m.CommandInput.Reset()
				return m, func() tea.Msg {
					return ModeNormalMsg(true)
				}
			}
		case "enter":
			if m.CommandInput.Focused() {
				cmd = m.handleCommand()
				cmds = append(cmds, cmd)
				m.CommandInput.Reset()
				m.CommandInput.Blur()
				m.Mode = Normal
			}
		}
	}

	m.CommandInput, cmd = m.CommandInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var ModeColor string
	var mode string
	switch m.Mode {
	case Insert:
		mode = "INSERT"
		ModeColor = "39"
	case Command:
		mode = "COMMAND"
		ModeColor = "70"
	default:
		mode = "NORMAL"
		ModeColor = "220"
	}
	modeStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#333333")).
		Width(9).
		Background(lipgloss.Color(ModeColor)).
		Bold(true).
		Padding(0, 1)
	commandStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#333333")).
		Width(tui.WindowSize.Width-9).
		Padding(0, 1)

	return lipgloss.JoinHorizontal(0, modeStyle.Render(mode), commandStyle.Render(m.CommandInput.View()))
}

func (m Model) handleCommand() tea.Cmd {
	command := m.CommandInput.Value()
	switch command {
	case ":q":
		return tea.Quit
	case ":g", ":generate":
		return func() tea.Msg {
			return SpawnGenerateMsg(true)
		}
	case ":d", ":deployments":
		return func() tea.Msg {
			return SpawnDeploymentsMsg(true)
		}
	}
	return nil
}
