package terminal

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/creack/pty"
)

func DetectDefaultShell() (string, []string) {
	switch runtime.GOOS {
	case "windows":
		if _, err := exec.LookPath("powershell.exe"); err == nil {
			return "powershell.exe", []string{"-NoLogo"}
		}

		return "cmd.exe", []string{}
	default:
		if sh := os.Getenv("SHELL"); sh != "" {
			return sh, []string{"-l"}
		}

		if _, err := exec.LookPath("bash"); err == nil {
			return "bash", []string{"-l"}
		}

		return "sh", []string{"-l"}
	}
}

func StartShell(cols, rows int) (io.WriteCloser, io.Reader, *exec.Cmd, error) {
	bin, args := DetectDefaultShell()
	cmd := exec.Command(bin, args...)

	cmd.Env = append(os.Environ(), "TERM=xterm-256color")

	if runtime.GOOS == "windows" {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return nil, nil, nil, err
		}

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, nil, nil, err
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return nil, nil, nil, err
		}

		merged := io.MultiReader(stdout, stderr)

		if err := cmd.Start(); err != nil {
			return nil, nil, nil, err
		}

		return stdin, merged, cmd, nil
	}

	ts := pty.Winsize{Cols: uint16(cols), Rows: uint16(rows)}

	ptmx, err := pty.StartWithSize(cmd, &ts)
	if err != nil {
		return nil, nil, nil, err
	}

	if cmd.Process == nil {
		ptmx.Close()

		return nil, nil, nil, errors.New("error when starting shell")
	}

	return ptmx, ptmx, cmd, nil
}
