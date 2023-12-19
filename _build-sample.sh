#!/bin/sh

## Static Build: https://github.com/golang/go/issues/26492
# windows: -tags netgo -ldflags '-H=windowsgui -extldflags "-static"'
# linux/bsd: -tags netgo -ldflags '-extldflags "-static"'
# macos: -ldflags '-s -extldflags "-sectcreate __TEXT __info_plist Info.plist"'
# android: -ldflags -s
##

# OpenGL Version can be modified in: internal/painter/gl/gl_core.go
# Default is 2.1
# alternate vendor files for opengl are in: .vendor-alt

go build -trimpath -mod=vendor -ldflags="-s -w" ./cmd/fyne_demo/

