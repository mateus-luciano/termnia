package core

import (
	"fmt"
	"io"
	"os/exec"
)

type BashTerminal struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser
}

func NewBashTerminal() *BashTerminal {
	return &BashTerminal{}
}

func (b *BashTerminal) Start() error {
	b.cmd = exec.Command("bash")

	var err error
	b.stdin, err = b.cmd.StdinPipe()
	if err != nil {
		return err
	}

	b.stdout, err = b.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	b.stderr, err = b.cmd.StderrPipe()
	if err != nil {
		return err
	}

	return b.cmd.Start()
}

func (b *BashTerminal) Stdin() io.Writer {
	return b.stdin
}

func (b *BashTerminal) Stdout() io.Reader {
	return b.stdout
}

func (b *BashTerminal) Stderr() io.Reader {
	return b.stderr
}

func (b *BashTerminal) Kill() error {
	if b.cmd.Process == nil {
		return fmt.Errorf("process not started")
	}

	return b.cmd.Process.Kill()
}
