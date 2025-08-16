package tui

func (m Model) View() string {
	s := "=== Termnia ===\n\n"

	start := 0
	if len(m.Output) > 15 {
		start = len(m.Output) - 15
	}

	for i := start; i < len(m.Output); i++ {
		s += m.Output[i] + "\n"
	}

	s += "\n> " + m.Input + "_"
	s += "\n\n[Ctrl+Q to exit]"
	return s
}
