//go:build ios
// +build ios

package mobile

import "github.com/unix-world/smart-fyne"

const tapYOffset = -12.0 // to compensate for how we hold our fingers on the device

func (*device) SystemScaleForWindow(_ fyne.Window) float32 {
	if currentDPI >= 450 {
		return 3
	} else if currentDPI >= 340 {
		return 2.5
	}
	return 2
}
