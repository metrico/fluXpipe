package main

import (
	"C"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/csv"
	_ "github.com/influxdata/flux/fluxinit/static"
	"github.com/influxdata/flux/lang"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/runtime"

	_fluxhttp "github.com/influxdata/flux/dependencies/http"
	"github.com/influxdata/flux/dependencies/secret"
	"github.com/influxdata/flux/dependencies/url"

	_ "embed"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SHARED
// # CGO_ENABLED=1 go build -buildmode=c-shared -o fluxpipe.a fluxpipelib.go

// STATIC
// # CGO_ENABLED=1 go build -buildmode=c-archive -o fluxpipe.a fluxpipelib.go

var APPNAME = "fluxpipe-library"

func runQuery(ctx context.Context, script string) (flux.Query, error) {

	program, err := lang.Compile(ctx, script, runtime.Default, time.Now())
	if err != nil {
		return nil, err
	}

	q, err := program.Start(ctx, memory.DefaultAllocator)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// NewCustomDependencies produces a Custom set of dependencies including EnvironmentSecretService.
func NewCustomDependencies() flux.Deps {
	validator := url.PassValidator{}
	return flux.Deps{
		Deps: flux.WrappedDeps{
			HTTPClient: _fluxhttp.NewLimitedDefaultClient(validator),
			// Default to having no filesystem, no secrets, and no url validation (always pass).
			FilesystemService: nil,
			SecretService:     secret.EnvironmentSecretService{},
			URLValidator:      validator,
		},
	}
}

func exec(inputString string) (string, error) {

	// CustomDeps produces a Custom set of dependencies including EnvironmentSecretService.
	customValidator := url.PassValidator{}
	customDeps := flux.Deps{
		Deps: flux.WrappedDeps{
			HTTPClient:        _fluxhttp.NewLimitedDefaultClient(customValidator),
			FilesystemService: nil,
			SecretService:     secret.EnvironmentSecretService{},
			URLValidator:      customValidator,
		},
	}

	// ctx := flux.NewDefaultDependencies().Inject(context.Background())
	ctx := customDeps.Inject(context.Background())

	q, err := runQuery(ctx, inputString)
	if err != nil {
		fmt.Println("unexpected error while creating query: %s", err)
		return "", err
	}

	results := flux.NewResultIteratorFromQuery(q)
	defer results.Release()

	buf := bytes.NewBuffer(nil)
	encoder := csv.NewMultiResultEncoder(csv.DefaultEncoderConfig())

	if _, err := encoder.Encode(buf, results); err != nil {
		return "", err
	}

	q.Done()

	if q.Err() != nil {
		fmt.Println("unexpected error from query execution: %s", q.Err())
		return "", q.Err()

	} else {
		return buf.String(), nil
	}
}

type FluxEvent struct {
	Query string `json:"flux"`
}

//export Query
func Query(query string) string {
	res, err := exec(query)
	if err != nil {
		return fmt.Sprintf(`{"code":"invalid","message":"%v"}`, err.Error())
	}
	return res
}

func main() {}
