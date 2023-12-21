//go:build !hints
// +build !hints

package theme

import (
	"image/color"

	"github.com/unix-world/smart-fyne"
)

var (
	fallbackColor = color.Transparent
	fallbackIcon  = &fyne.StaticResource{}
)
