#!/bin/sh

go build -trimpath -mod=vendor -ldflags="-s -w" ./cmd/fyne_demo/

