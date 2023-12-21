//go:build android
// +build android

package mobile

/*
#cgo LDFLAGS: -landroid -llog -lEGL -lGLESv2

#include <stdbool.h>
#include <stdlib.h>

bool canListURI(uintptr_t jni_env, uintptr_t ctx, char* uriCstr);
bool createListableURI(uintptr_t jni_env, uintptr_t ctx, char* uriCstr);
char *listURI(uintptr_t jni_env, uintptr_t ctx, char* uriCstr);
*/
import "C"
import (
	"errors"
	"strings"
	"unsafe"

	"github.com/unix-world/smart-fyne"
	"github.com/unix-world/smart-fyne/internal/driver/mobile/app"
	"github.com/unix-world/smart-fyne/storage"
)

func canListURI(uri fyne.URI) bool {
	uriStr := C.CString(uri.String())
	defer C.free(unsafe.Pointer(uriStr))
	listable := false

	app.RunOnJVM(func(_, env, ctx uintptr) error {
		listable = bool(C.canListURI(C.uintptr_t(env), C.uintptr_t(ctx), uriStr))
		return nil
	})

	return listable
}

func createListableURI(uri fyne.URI) error {
	uriStr := C.CString(uri.String())
	defer C.free(unsafe.Pointer(uriStr))

	ok := false
	app.RunOnJVM(func(_, env, ctx uintptr) error {
		ok = bool(C.createListableURI(C.uintptr_t(env), C.uintptr_t(ctx), uriStr))
		return nil
	})

	if ok {
		return nil
	}
	return errors.New("failed to create directory")
}

func listURI(uri fyne.URI) ([]fyne.URI, error) {
	uriStr := C.CString(uri.String())
	defer C.free(unsafe.Pointer(uriStr))

	var str *C.char
	app.RunOnJVM(func(_, env, ctx uintptr) error {
		str = C.listURI(C.uintptr_t(env), C.uintptr_t(ctx), uriStr)
		return nil
	})

	parts := strings.Split(C.GoString(str), "|")
	var list []fyne.URI
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		list = append(list, storage.NewURI(part))
	}
	return list, nil
}
