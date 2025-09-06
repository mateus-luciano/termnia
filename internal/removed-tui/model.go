package removedUi

// AMAZON Q IGNORAR ESSE ARQUIVO

import (
	"time"

	"github.com/mateus-luciano/termnia/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

type Tab struct {
	Shell  *core.ShellTerminal
	Output []string
	Input  string
	Name   string
}

type Model struct {
	Tabs       []Tab
	ActiveTab  int
	TabRegions []TabRegion
	NewTabBtn  ButtonRegion
}

type TabRegion struct {
	StartX, EndX int
	TabIndex     int
}

type ButtonRegion struct {
	StartX, EndX int
	Y            int
}

func NewModel(shell *core.ShellTerminal) Model {
	firstTab := Tab{
		Shell:  shell,
		Output: []string{"Termnia v1.0.0", "Digite comandos abaixo:"},
		Input:  "",
		Name:   "Terminal 1",
	}

	return Model{
		Tabs:       []Tab{firstTab},
		ActiveTab:  0,
		TabRegions: []TabRegion{},
		NewTabBtn:  ButtonRegion{},
	}
}

func (m Model) Init() tea.Cmd {
	if len(m.Tabs) > 0 && m.Tabs[m.ActiveTab].Shell != nil {
		return readShellOutput(m.Tabs[m.ActiveTab].Shell)
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
