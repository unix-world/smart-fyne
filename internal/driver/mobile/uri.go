//go:build !android
// +build !android

package mobile

import (
	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/storage"
)

func nativeURI(path string) fyne.URI {
	uri, err := storage.ParseURI(path)
	if err != nil {
		fyne.LogError("Error on parsing uri", err)
	}
	return uri
}
