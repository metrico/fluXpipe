#!/bin/sh

echo "Building fluxpipe..."

export PKG_CONFIG_PATH=$(pwd)

# go build -a -ldflags '-extldflags "-static -w -ldl"' -o fluxpipe fluxpipe.go
go build -a -ldflags '-extldflags "-static -w -ldl"' -o fluxpipe-server fluxpipe-server.go
