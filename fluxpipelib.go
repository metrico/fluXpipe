package main

import (
	"C"
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/InfluxCommunity/flux"
	"github.com/InfluxCommunity/flux/csv"
	_ "github.com/InfluxCommunity/flux/fluxinit/static"
	"github.com/InfluxCommunity/flux/lang"
	"github.com/InfluxCommunity/flux/memory"
	"github.com/InfluxCommunity/flux/runtime"

	_fluxhttp "github.com/InfluxCommunity/flux/dependencies/http"
	"github.com/InfluxCommunity/flux/dependencies/secret"
	"github.com/InfluxCommunity/flux/dependencies/url"
)


// # CGO_ENABLED=1 go build -buildmode=c-shared -o fluxpipe.a fluxpipelib.go
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
