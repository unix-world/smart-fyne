//go:build !ci && !js && !wasm && test_web_driver
// +build !ci,!js,!wasm,test_web_driver

package app

import (
	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/theme"
)

func defaultVariant() fyne.ThemeVariant {
	return theme.VariantDark
}
