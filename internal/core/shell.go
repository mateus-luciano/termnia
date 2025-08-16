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
		return nil, fmt.Errorf("shell '%s' is not available on this system", shell)
	}
	return &ShellTerminal{name: shell, cmd: exec.Command(cmdPath)}, nil
}

func DefaultShellForOS() ShellType {
	switch runtime.GOOS {
	case "windows":
		return ShellCmd
	case "linux", "darwin":
		return ShellBash
	default:
		return ShellBash
	}
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
