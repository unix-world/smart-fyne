// Package desktop provides desktop specific driver functionality.
package desktop

import "github.com/unix-world/smart-fyne"

// Driver represents the extended capabilities of a desktop driver
type Driver interface {
	// Create a new borderless window that is centered on screen
	CreateSplashWindow() fyne.Window

	// Gets the set of key modifiers that are currently active
	//
	// Since: 2.4
	CurrentKeyModifiers() fyne.KeyModifier
}
