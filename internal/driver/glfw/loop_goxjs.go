//go:build js || wasm || test_web_driver
// +build js wasm test_web_driver

package glfw

import (
	"fmt"

	"github.com/unix-world/smart-fyne"

	gl "github.com/unix-world/smart-fyne/gl-js"
	glfw "github.com/unix-world/smart-fyne/glfw-js"
)

func (d *gLDriver) initGLFW() {
	initOnce.Do(func() {
		err := glfw.Init(gl.ContextWatcher)
		if err != nil {
			fyne.LogError("failed to initialise GLFW", err)
			return
		}

		d.startDrawThread()
	})
}

func (d *gLDriver) tryPollEvents() {
	defer func() {
		if r := recover(); r != nil {
			fyne.LogError(fmt.Sprint("GLFW poll event error: ", r), nil)
		}
	}()

	glfw.PollEvents() // This call blocks while window is being resized, which prevents freeDirtyTextures from being called
}

func (d *gLDriver) Terminate() {
	glfw.Terminate()
}
