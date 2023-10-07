#!/bin/sh

# WARNING: 
# Make sure you have a complete flux build at `/usr/src/flux`

export PKG_CONFIG_PATH=$(pwd)

echo "Building fluxpipe-server ..."
go build -a -ldflags '-extldflags "-static -w -ldl"' -o fluxpipe-server ./cmd/server

echo "Building fluxpipe-lambda ..."
go build -a -ldflags '-extldflags "-static -w -ldl"' -o fluxpipe-lambda ./cmd/lambda

echo "Building fluxpipe-lib ..."
CGO_ENABLED=1 go build -buildmode=c-archive -o fluxpipelib.a ./cmd/lib
CGO_ENABLED=1 go tool cgo -exportheader ./libfluxpipe.h ./cmd/lib/fluxpipelib.go
# CGO_ENABLED=1 go build -buildmode=c-shared -o fluxpipelib.dylib fluxpipelib.go
