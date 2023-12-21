//go:build !ios && !android && !wayland
// +build !ios,!android,!wayland

package mobile

import "github.com/unix-world/smart-fyne"

const tapYOffset = 0 // no finger compensation on desktop (simulation)

func (*device) SystemScaleForWindow(_ fyne.Window) float32 {
	return 2 // this is simply due to the high number of pixels on a mobile device - just an approximation
}
