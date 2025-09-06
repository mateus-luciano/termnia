package core

import (
	"fmt"
	"io"
	"os/exec"
	"runtime"

	"github.com/mateus-luciano/termnia/internal/platform"
	"github.com/mateus-luciano/termnia/internal/types"
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
	name   types.ShellType
}

func NewShellTerminal(shell types.ShellType) (*ShellTerminal, error) {
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

	cmd := exec.Command(cmdPath)
	platform.ConfigurePlatformSpecific(cmd)

	return &ShellTerminal{name: shell, cmd: cmd}, nil
}

func DefaultShellForOS() types.ShellType {
	var priority []types.ShellType

	switch runtime.GOOS {
	case "windows":
		priority = []types.ShellType{types.ShellCmd, types.ShellPowerShell, types.ShellWSL}
	case "linux", "darwin":
		priority = []types.ShellType{types.ShellBash, types.ShellZsh}
	default:
		priority = []types.ShellType{types.ShellBash, types.ShellZsh}
	}

	for _, shell := range priority {
		if detectShell(shell) != "" {
			return shell
		}
	}

	return priority[0]
}

func detectShell(shell types.ShellType) string {
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

func (s *ShellTerminal) Stdin() io.Writer {
	return s.stdin
}

func (s *ShellTerminal) Stdout() io.Reader {
	return s.stdout
}

func (s *ShellTerminal) Stderr() io.Reader {
	return s.stderr
}

func (s *ShellTerminal) Name() types.ShellType {
	return s.name
}

func (s *ShellTerminal) Kill() error {
	if s.cmd.Process == nil {
		return fmt.Errorf("process not started")
	}

	return s.cmd.Process.Kill()
}
