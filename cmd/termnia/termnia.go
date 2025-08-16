package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"termnia/internal/core"
)

func main() {
	shell, err := core.NewShellTerminal(defaultShellForOS())
	if err != nil {
		panic(err)
	}

	if err := shell.Start(); err != nil {
		panic(err)
	}

	go func() {
		scanner := bufio.NewScanner(shell.Stdout())
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text() + "\n"
		shell.Stdin().Write([]byte(line))
	}
}

func defaultShellForOS() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "cmd"
	case "linux", "darwin":
		return "bash"
	default:
		return "bash"
	}
}
