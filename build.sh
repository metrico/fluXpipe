#!/bin/sh

# WARNING: 
# Make sure you have a complete flux build at `/usr/src/flux`

export PKG_CONFIG_PATH=$(pwd)

echo "Building fluxpipe-server ..."
go build -a -ldflags '-extldflags "-static -w -ldl"' -o fluxpipe-server fluxpipe-server.go

echo "Building fluxpipe-lambda ..."
go build -a -ldflags '-extldflags "-static -w -ldl"' -o fluxpipe-lambda fluxpipe-lambda.go

echo "Building fluxpipe-lib ..."
CGO_ENABLED=1 go build -buildmode=c-archive -o fluxpipelib.a fluxpipelib.go
# CGO_ENABLED=1 go build -buildmode=c-shared -o fluxpipelib.dylib fluxpipelib.go
