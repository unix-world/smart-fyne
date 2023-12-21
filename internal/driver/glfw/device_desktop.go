//go:build !js && !wasm
// +build !js,!wasm

package glfw

import (
	"runtime"

	"github.com/unix-world/smart-fyne"
)

func (*glDevice) IsMobile() bool {
	return false
}

func (*glDevice) SystemScaleForWindow(w fyne.Window) float32 {
	if runtime.GOOS == "darwin" {
		return 1.0 // macOS scaling is done at the texture level
	}
	if runtime.GOOS == "windows" {
		xScale, _ := w.(*window).viewport.GetContentScale()
		return xScale
	}

	return scaleAuto
}
