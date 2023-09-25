package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func stepCreateView(m lmay) string {
	var b strings.Builder

	writeSpacedText(&b, "Welcome to the step creation view.")
	writeSpacedText(&b, "Here are your matched files:")
	for _, v := range m.matchedFiles {
		writeSpacedText(&b, fmt.Sprintf("%s", v))
	}

	writeHelpFooter(&b)

	return b.String()
}

func updateStepCreation(msg tea.Msg, m lmay) (lmay, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			m.stage = confirmation
			return m, cmd
		}
	}

	return m, cmd
}
