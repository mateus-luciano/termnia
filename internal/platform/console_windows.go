package platform

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	procAllocConsole    = kernel32.NewProc("AllocConsole")
	procSetConsoleTitle = kernel32.NewProc("SetConsoleTitleW")
)

func AllocConsole() error {
	ret, _, err := procAllocConsole.Call()
	if ret == 0 {
		return err
	}

	title, _ := syscall.UTF16PtrFromString("Termnia")
	procSetConsoleTitle.Call(uintptr(unsafe.Pointer(title)))

	return nil
}

func RedirectIO() error {
	stdin, err := os.OpenFile("CONIN$", os.O_RDWR, 0)
	if err != nil {
		return err
	}

	stdout, err := os.OpenFile("CONOUT$", os.O_RDWR, 0)
	if err != nil {
		return err
	}

	stderr, err := os.OpenFile("CONOUT$", os.O_RDWR, 0)
	if err != nil {
		return err
	}

	os.Stdin = stdin
	os.Stdout = stdout
	os.Stderr = stderr

	return nil
}
