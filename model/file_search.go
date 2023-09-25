package model

import (
	"io/fs"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func fileSearchView(m lmay) string {
	var b strings.Builder

	writeSpacedText(&b, focusedStyle.Render("Searching for matching files..."))

	b.WriteRune('\n')
	writeSpacedText(&b, helpStyle.Render("(shift+tab) to go back | (esc) to quit"))

	return b.String()
}

func updateFileSearch(msg tea.Msg, m lmay) (lmay, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyShiftTab:
			m.matchedFiles = nil

			m.textInput.SetValue("")
			m.stage = fileSelect

			return m, cmd
		}
	}

	m, cmd = updateMatchedFiles(m)

	return m, cmd
}

func updateMatchedFiles(m lmay) (lmay, tea.Cmd) {
	var cmd tea.Cmd

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && m.fileRegex.Match([]byte(d.Name())) {
			m.matchedFiles = append(m.matchedFiles, path)
		}

		return nil
	})

	if err != nil {
		cmd = returnErrMsg(m, err)
		return m, cmd
	}

	m.stage = stepCreation

	return m, cmd
}
