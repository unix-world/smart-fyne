package test

import (
	"os"

	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/internal"
	"github.com/unix-world/smart-fyne/storage"
)

type testStorage struct {
	*internal.Docs
}

func (s *testStorage) RootURI() fyne.URI {
	return storage.NewFileURI(os.TempDir())
}

func (s *testStorage) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
