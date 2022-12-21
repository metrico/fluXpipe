#!/bin/sh
export PKG_CONFIG_PATH=$(pwd)

echo "Building fluxpipe-server ..."
go mod tidy
go build -a -ldflags '-extldflags "-static -w -ldl"' -o fluxpipe-server fluxpipe-server.go
