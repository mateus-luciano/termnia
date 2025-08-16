package core

import (
	"fmt"
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
	name   ShellType
}

func NewShellTerminal(shell ShellType) (*ShellTerminal, error) {
	cmdPath := detectShell(shell)
	if cmdPath == "" {
		fallback := DefaultShellForOS()
		if fallback == shell {
			return nil, fmt.Errorf("shell '%s' is not available on this system", shell)
		}
		shell = fallback
		cmdPath = detectShell(shell)
		if cmdPath == "" {
			return nil, fmt.Errorf("no available shells found on this system")
		}
	}
	return &ShellTerminal{name: shell, cmd: exec.Command(cmdPath)}, nil
}

func DefaultShellForOS() ShellType {
	var priority []ShellType
	switch runtime.GOOS {
	case "windows":
		priority = []ShellType{ShellCmd, ShellPowerShell, ShellWSL}
	case "linux", "darwin":
		priority = []ShellType{ShellBash, ShellZsh}
	default:
		priority = []ShellType{ShellBash, ShellZsh}
	}

	for _, shell := range priority {
		if detectShell(shell) != "" {
			return shell
		}
	}
	return priority[0]
}

func detectShell(shell ShellType) string {
	os := runtime.GOOS
	if paths, ok := ShellPaths[os]; ok {
		if path, ok := paths[shell]; ok {
			if _, err := exec.LookPath(path); err == nil {
				return path
			}
		}
	}
	return ""
}

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
func (s *ShellTerminal) Name() ShellType   { return s.name }
func (s *ShellTerminal) Kill() error       { return s.cmd.Process.Kill() }
