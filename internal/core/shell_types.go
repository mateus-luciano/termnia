package core

type ShellType string

const (
	ShellCmd        ShellType = "cmd"
	ShellPowerShell ShellType = "powershell"
	ShellWSL        ShellType = "wsl"
	ShellBash       ShellType = "bash"
	ShellZsh        ShellType = "zsh"
)
