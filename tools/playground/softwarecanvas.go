package playground

import (
	"github.com/unix-world/smart-fyne/driver/software"
	"github.com/unix-world/smart-fyne/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return software.NewCanvas()
}
