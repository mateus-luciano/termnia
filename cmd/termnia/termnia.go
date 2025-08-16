package main

import (
	"os"
	"termnia/internal/config"
	"termnia/internal/core"
	"termnia/internal/platform"
	"termnia/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if err := platform.AllocConsole(); err != nil {
		os.Exit(1)
	}
	if err := platform.RedirectIO(); err != nil {
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		os.Exit(1)
	}

	shell, err := core.NewShellTerminal(cfg.DefaultShell)
	if err != nil || shell == nil {
		shell, _ = core.NewShellTerminal(core.DefaultShellForOS())
	}
	if err := shell.Start(); err != nil {
		os.Exit(1)
	}

	model := tui.NewModel(shell)
	p := tea.NewProgram(model)
	_, _ = p.Run()
}
