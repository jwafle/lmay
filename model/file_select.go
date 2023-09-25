package model

import (
	"fmt"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	validFileExtensions = []string{".go", ".yaml", ".yml"}
)

func validateFileExtension(filename string) error {
	for _, ext := range validFileExtensions {
		if strings.HasSuffix(filename, ext) {
			return nil
		}
	}
	return fmt.Errorf("file extension not supported for file: %s", filename)
}

func generateRegex(filename string) (*regexp.Regexp, error) {
	var b strings.Builder
	b.WriteString(`\b`)
	for _, char := range filename {
		switch char {
		case '.':
			b.WriteString(`\.`)
		case '*':
			b.WriteString(`.*`)
		default:
			b.WriteRune(char)
		}
	}
	b.WriteString(`\b`)

	return regexp.Compile(b.String())
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

			regex, err := generateRegex(m.textInput.Value())
			if err != nil {
				cmd = returnErrMsg(m, err)
				return m, cmd
			}

			m.fileRegex = regex
			m.stage = fileSearch
			return m, cmd
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
