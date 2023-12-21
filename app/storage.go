package app

import (
	"os"

	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/internal"
	"github.com/unix-world/smart-fyne/storage"
)

type store struct {
	*internal.Docs
	a *fyneApp
}

func (s *store) RootURI() fyne.URI {
	if s.a.UniqueID() == "" {
		fyne.LogError("Storage API requires a unique ID, use app.NewWithID()", nil)
		return storage.NewFileURI(os.TempDir())
	}

	return storage.NewFileURI(s.a.storageRoot())
}

func (s *store) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
