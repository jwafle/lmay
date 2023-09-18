package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func userErrorView(m lmay) string {
	var b strings.Builder

	writeSpacedText(&b, errorStyle.Render("Sorry, an error appears to have occurred."))

	b.WriteRune('\n')
	b.WriteString(errorStyle.Render("Error: "))
	b.WriteString(standardStyle.Render(m.err.Error()))
	b.WriteRune('\n')

	writeHelpFooter(&b)

	return b.String()
}

func updateUserError(msg tea.Msg, m lmay) (lmay, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			m.err = nil
			m.stage = fileSelect
			return m, cmd
		}
	}

	return m, cmd
}
