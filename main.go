package main

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"log/slog"
	"os"

	"github.com/jwafle/lmay/model"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

func main() {
	if _, err := tea.NewProgram(model.InitialModel()).Run(); err != nil {
		logger.Error("failed to init program", slog.String("error", err.Error()))
	}
}

type ()
