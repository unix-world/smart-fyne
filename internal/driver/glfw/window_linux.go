package glfw

import "github.com/unix-world/smart-fyne"

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
}
