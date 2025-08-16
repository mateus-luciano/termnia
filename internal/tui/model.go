package tui

import (
	"termnia/internal/core"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Shell  *core.ShellTerminal
	Output []string
	Input  string
}

func NewModel(shell *core.ShellTerminal) Model {
	return Model{
		Shell:  shell,
		Output: []string{"Enter commands below:"},
		Input:  "",
	}
}

func (m Model) Init() tea.Cmd {
	if m.Shell != nil {
		return readShellOutput(m.Shell)
	}
	return nil
}

type shellOutputMsg struct {
	output string
}

func readShellOutput(shell *core.ShellTerminal) tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		buf := make([]byte, 1024)
		n, err := shell.Stdout().Read(buf)
		if err == nil && n > 0 {
			return shellOutputMsg{output: string(buf[:n])}
		}
		return nil
	})
}
