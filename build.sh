#!/bin/sh

GOARCH=amd64

GOOS=linux
go build -o aview-linux

if [ "$1" = "-a" ] || [ "$1" = "--all" ]; then
GOOS=windows
go build

GOOS=darwin
go build -o aview-mac
fi