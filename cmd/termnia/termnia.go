package main

import (
	"bufio"
	"fmt"
	"os"
	"termnia/internal/config"
	"termnia/internal/core"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	shell, err := core.NewShellTerminal(cfg.DefaultShell)
	if err != nil || shell == nil {
		shell, _ = core.NewShellTerminal(core.DefaultShellForOS())
	}
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
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading stdout: %v\n", err)
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text() + "\n"
		shell.Stdin().Write([]byte(line))
	}
}
