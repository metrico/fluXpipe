package main

import (
	"C"
	"context"
	"fmt"
	_ "github.com/InfluxCommunity/flux/fluxinit/static"
	"github.com/metrico/fluxpipe/service"
)

// # CGO_ENABLED=1 go build -buildmode=c-shared -o fluxpipe.a fluxpipelib.go
// # CGO_ENABLED=1 go build -buildmode=c-archive -o fluxpipe.a fluxpipelib.go

var APPNAME = "fluxpipe-library"

//export Query
func Query(query string) string {
	res, err := service.RunE(context.Background(), query)
	if err != nil {
		return fmt.Sprintf(`{"code":"invalid","message":"%v"}`, err.Error())
	}
	return res
}

func main() {}
