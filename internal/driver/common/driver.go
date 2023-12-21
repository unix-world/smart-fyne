package common

import (
	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/internal/cache"
)

// CanvasForObject returns the canvas for the specified object.
func CanvasForObject(obj fyne.CanvasObject) fyne.Canvas {
	return cache.GetCanvasForObject(obj)
}
