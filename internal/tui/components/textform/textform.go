package textform

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gregmulvaney/bosun/internal/tui"
	"github.com/gregmulvaney/bosun/internal/tui/components/button"
	"strings"
)

var textformStyle = lipgloss.NewStyle()

type CompleteMsg bool

type Model struct {
	focusIndex int
	prompt     string
	inputs     []textinput.Model
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
		t.PromptStyle = tui.PrimaryColor
		switch i {
		case 0:
			t.Focus()
		}
		inputs[i] = t
	}

	m := Model{
		focusIndex: 0,
		prompt:     prompt,
		inputs:     inputs,
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
			if m.focusIndex == len(m.inputs) {

			}
		case "up", "down", "tab":
			s := msg.String()
			if s == "up" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}
			if m.focusIndex < 0 {
				m.focusIndex = 0
			} else if m.focusIndex > len(m.inputs) {
				m.focusIndex = len(m.inputs)
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i < len(m.inputs); i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				m.inputs[i].Blur()
			}
			return m, tea.Batch(cmds...)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString("\n" + m.prompt + "\n\n")
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View() + "\n")
	}
	btn := button.New("Next", m.focusIndex, len(m.inputs))
	b.WriteString("\n" + btn.View() + "\n")
	return textformStyle.Width(m.width).Render(b.String())
}
