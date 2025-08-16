package core

var ShellPaths = map[string]map[string]string{
	"windows": {
		"cmd":        "cmd.exe",
		"powershell": "powershell.exe",
		"wsl":        "wsl.exe",
	},
	"darwin": {
		"bash": "/bin/bash",
		"zsh":  "/bin/zsh",
	},
	"linux": {
		"bash": "/bin/bash",
		"zsh":  "/bin/zsh",
	},
}
