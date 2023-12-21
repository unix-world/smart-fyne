//go:build !ci && (android || ios || mobile)
// +build !ci
// +build android ios mobile

package app

import (
	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/internal/driver/mobile"
)

var systemTheme fyne.ThemeVariant

// NewWithID returns a new app instance using the appropriate runtime driver.
// The ID string should be globally unique to this app.
func NewWithID(id string) fyne.App {
	d := mobile.NewGoMobileDriver()
	a := newAppWithDriver(d, id)
	d.(mobile.ConfiguredDriver).SetOnConfigurationChanged(func(c *mobile.Configuration) {
		systemTheme = c.SystemTheme

		a.Settings().(*settings).setupTheme()
	})
	return a
}
