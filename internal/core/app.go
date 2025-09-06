package core

import (
	"os"

	"github.com/mateus-luciano/termnia/internal/config"
	"github.com/mateus-luciano/termnia/internal/platform"
	"github.com/mateus-luciano/termnia/internal/ui"
)

func Run() error {
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

	shell, err := NewShellTerminal(cfg.DefaultShell)
	if err != nil || shell == nil {
		shell, _ = NewShellTerminal(DefaultShellForOS())
	}

	if err := shell.Start(); err != nil {
		os.Exit(1)
	}

	return ui.Run()
}
