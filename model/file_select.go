package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	validFileExtensions = []string{".yaml", ".yml"}
)

func validateFileExtension(filename string) error {
	for _, ext := range validFileExtensions {
		if strings.HasSuffix(filename, ext) {
			return nil
		}
	}
	return fmt.Errorf("file extension not supported for file: %s", filename)
}

func fileSelectView(m lmay) string {
	var b strings.Builder

	writeSpacedText(&b, focusedStyle.Render("What file(s) are you looking for?"))

	writeSpacedText(&b, fmt.Sprintf(
		"\n%s\n",
		m.textInput.View(),
	))

	writeHelpFooter(&b)

	return b.String()
}

func updateFileSelect(msg tea.Msg, m lmay) (lmay, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			err := validateFileExtension(m.textInput.Value())

			if err != nil {
				cmd = returnErrMsg(m, err)
				return m, cmd
			}

			m.stage = stepCreation
			return m, cmd
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
