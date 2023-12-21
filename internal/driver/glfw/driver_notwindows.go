//go:build !windows
// +build !windows

package glfw

import "github.com/unix-world/smart-fyne"

func logError(msg string, err error) {
	fyne.LogError(msg, err)
}

func isDark() bool {
	return true // this is really a no-op placeholder for a windows menu workaround
}
