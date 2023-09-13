package main

import (
	"context"
	"fmt"
	"time"
	"bytes"

	"github.com/InfluxCommunity/flux"
	"github.com/InfluxCommunity/flux/csv"
	_ "github.com/InfluxCommunity/flux/fluxinit/static"
	"github.com/InfluxCommunity/flux/lang"
	"github.com/InfluxCommunity/flux/memory"
	"github.com/InfluxCommunity/flux/runtime"

	_fluxhttp "github.com/InfluxCommunity/flux/dependencies/http"
	"github.com/InfluxCommunity/flux/dependencies/secret"
	"github.com/InfluxCommunity/flux/dependencies/url"
  
	"github.com/aws/aws-lambda-go/lambda"

)

var APPNAME = "fluxpipe-lambda"

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

func exec(inputString string) (string, string) {

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
		return "", string(fmt.Sprintf(`{"code":"invalid","message":"%v"}`, err))
	}

	results := flux.NewResultIteratorFromQuery(q)
	defer results.Release()

	buf := bytes.NewBuffer(nil)
	encoder := csv.NewMultiResultEncoder(csv.DefaultEncoderConfig())

	if _, err := encoder.Encode(buf, results); err != nil {
		panic(err)
	}

	q.Done()

	if q.Err() != nil {
		fmt.Println("unexpected error from query execution: %s", q.Err())
		return "", string(fmt.Sprintf(`{"code":"invalid","message":"%v"}`, q.Err()))

	} else {
		return buf.String(), ""
	}
}

type FluxEvent struct {
        Query string `json:"flux"`
}

func HandleRequest(ctx context.Context, flux FluxEvent) (string, error) {
        buf, _ := exec(flux.Query)
        return buf, nil
}

func main() {
        lambda.Start(HandleRequest)
}
