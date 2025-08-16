package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return m, tea.Quit
		case "enter":
			if m.Input != "" {
				m.Output = append(m.Output, "> "+m.Input)
				if m.Shell != nil {
					m.Shell.Stdin().Write([]byte(m.Input + "\n"))
				} else {
					m.Output = append(m.Output, "Shell not available")
				}
				m.Input = ""
			}
		case "backspace":
			if len(m.Input) > 0 {
				m.Input = m.Input[:len(m.Input)-1]
			}
		default:
			if len(msg.String()) == 1 {
				m.Input += msg.String()
			}
		}
	case shellOutputMsg:
		if msg.output != "" {
			m.Output = append(m.Output, msg.output)
		}
		return m, readShellOutput(m.Shell)
	}
	return m, nil
}
