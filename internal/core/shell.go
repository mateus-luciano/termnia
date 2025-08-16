package core

import (
	"errors"
	"io"
	"os/exec"
	"runtime"
)

type Terminal interface {
	Start() error
	Stdin() io.Writer
	Stdout() io.Reader
	Stderr() io.Reader
	Kill() error
}

type ShellTerminal struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser
	name   string
}

func NewShellTerminal(shell string) (*ShellTerminal, error) {
	cmdPath := detectShell(shell)
	if cmdPath == "" {
		return nil, errors.New("shell is not available on this system")
	}

	return &ShellTerminal{name: shell, cmd: exec.Command(cmdPath)}, nil
}

func detectShell(shell string) string {
	switch runtime.GOOS {
	case "windows":
		switch shell {
		case "cmd":
			return "cmd.exe"
		case "powershell":
			return "powershell.exe"
		case "wsl":
			if path, err := exec.LookPath("wsl.exe"); err == nil {
				return path
			}
			return ""
		default:
			return "cmd.exe"
		}
	case "darwin":
		switch shell {
		case "bash":
			return "/bin/bash"
		case "zsh":
			return "/bin/zsh"
		default:
			return "/bin/bash"
		}
	case "linux":
		switch shell {
		case "bash":
			return "/bin/bash"
		case "zsh":
			return "/bin/zsh"
		default:
			return "/bin/bash"
		}
	default:
		return ""
	}
}

func (s *ShellTerminal) Name() string { return s.name }
func (s *ShellTerminal) Start() error {
	var err error
	s.stdin, err = s.cmd.StdinPipe()
	if err != nil {
		return err
	}
	s.stdout, err = s.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	s.stderr, err = s.cmd.StderrPipe()
	if err != nil {
		return err
	}
	return s.cmd.Start()
}
func (s *ShellTerminal) Stdin() io.Writer  { return s.stdin }
func (s *ShellTerminal) Stdout() io.Reader { return s.stdout }
func (s *ShellTerminal) Stderr() io.Reader { return s.stderr }
func (s *ShellTerminal) Kill() error       { return s.cmd.Process.Kill() }
