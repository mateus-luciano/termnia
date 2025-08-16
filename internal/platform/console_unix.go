//go:build !windows

package platform

func AllocConsole() error {
	return nil
}

func RedirectIO() error {
	return nil
}