package core

import "github.com/mateus-luciano/termnia/internal/types"

var ShellPaths = map[string]map[types.ShellType]string{
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
