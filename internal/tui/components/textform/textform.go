package textform

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui/components/button"
	"strings"
)

var textformStyle = lipgloss.NewStyle()

type CompleteMsg bool

type Model struct {
	FocusIndex int
	prompt     string
	Inputs     []textinput.Model
	fieldNames []string
	width      int
}

func New(fieldNames []string, prompt string) Model {
	inputs := make([]textinput.Model, len(fieldNames))
	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		t.Prompt = fieldNames[i] + ": "
		t.CharLimit = 64
		t.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
		inputs[i] = t
	}

	m := Model{
		FocusIndex: 0,
		prompt:     prompt,
		Inputs:     inputs,
		fieldNames: fieldNames,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ", "enter":
			if m.FocusIndex == len(m.Inputs) {

			}
		case "up", "down", "tab":
			s := msg.String()
			if s == "up" {
				m.FocusIndex--
			} else {
				m.FocusIndex++
			}
			if m.FocusIndex < 0 {
				m.FocusIndex = 0
			} else if m.FocusIndex > len(m.Inputs) {
				m.FocusIndex = len(m.Inputs)
			}
			cmds := make([]tea.Cmd, len(m.Inputs))
			for i := 0; i < len(m.Inputs); i++ {
				if i == m.FocusIndex {
					cmds[i] = m.Inputs[i].Focus()
					continue
				}
				m.Inputs[i].Blur()
			}
			return m, tea.Batch(cmds...)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	cmds := make([]tea.Cmd, len(m.Inputs))
	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString("\n" + m.prompt + "\n\n")
	for i := range m.Inputs {
		b.WriteString(m.Inputs[i].View() + "\n")
	}
	btn := button.New("Next", m.FocusIndex, len(m.Inputs))
	b.WriteString("\n" + btn.View() + "\n")
	return textformStyle.Width(m.width).Render(b.String())
}
