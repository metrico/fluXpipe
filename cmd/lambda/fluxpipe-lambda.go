package main

import (
	"context"
	"github.com/metrico/fluxpipe/service"

	"github.com/aws/aws-lambda-go/lambda"
)

var APPNAME = "fluxpipe-lambda"

func exec(ctx context.Context, inputString string) (string, string) {
	res, err := service.RunE(ctx, inputString)
	return res, err.Error()
}

type FluxEvent struct {
	Query string `json:"flux"`
}

func HandleRequest(ctx context.Context, flux FluxEvent) (string, error) {
	buf, _ := exec(ctx, flux.Query)
	return buf, nil
}

func main() {
	lambda.Start(HandleRequest)
}
