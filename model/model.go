package model

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	Stage  uint8
	errMsg error
)

const (
	userError Stage = iota
	fileSelect
	stepCreation
	confirmation
	execution
)

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	standardStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	helpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
)

type lmay struct {
	textinput textinput.Model
	stage     Stage
	err       error
}

func InitialModel() lmay {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.TextStyle = standardStyle

	return lmay{
		textinput: ti,
		stage:     fileSelect,
		err:       nil,
	}
}

func (m lmay) Init() tea.Cmd {
	return textinput.Blink
}

func (m lmay) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		m.stage = userError
		return m, cmd
	default:
		switch m.stage {
		case userError:
			m, cmd = updateUserError(msg, m)
		case fileSelect:
			m, cmd = updateFileSelect(msg, m)
			return m, cmd
		case stepCreation:
			m, cmd = updateStepCreation(msg, m)
			return m, cmd
		}
	}

	return m, cmd
}

func (m lmay) View() string {
	var s string

	switch m.stage {
	case userError:
		s = userErrorView(m)
	case fileSelect:
		s = fileSelectView(m)
	case stepCreation:
		s = stepCreateView(m)
	}

	return s
}

func returnErrMsg(l lmay, err error) tea.Cmd {
	return func() tea.Msg {
		return errMsg(err)
	}
}

func writeSpacedText(b *strings.Builder, s string) {
	b.WriteString(s)
	b.WriteRune('\n')
}

func writeHelpFooter(b *strings.Builder) {
	b.WriteRune('\n')
	writeSpacedText(b, helpStyle.Render("(esc) to quit"))
}
