package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/InfluxCommunity/flux"
	"github.com/InfluxCommunity/flux/csv"
	"github.com/InfluxCommunity/flux/dependencies/bigtable"
	"github.com/InfluxCommunity/flux/dependencies/influxdb"
	"github.com/InfluxCommunity/flux/dependencies/mqtt"
	"github.com/InfluxCommunity/flux/dependency"
	_ "github.com/InfluxCommunity/flux/fluxinit/static"
	"github.com/InfluxCommunity/flux/lang"
	"github.com/InfluxCommunity/flux/memory"
	"github.com/InfluxCommunity/flux/runtime"
)

func RunE(ctx context.Context, script string) (string, error) {
	// Defer initialization until other common errors
	// have already passed to avoid a long load time
	// for a simple unrelated error.

	ctx = injectDependencies(ctx)

	return executeE(ctx, script)
}

func executeE(ctx context.Context, script string) (string, error) {
	c := lang.FluxCompiler{
		Query: script,
	}
	prog, err := c.Compile(ctx, runtime.Default)
	if err != nil {
		return "", err
	}

	mem := &memory.ResourceAllocator{}
	q, err := prog.Start(ctx, mem)
	if err != nil {
		return "", err
	}

	results := flux.NewResultIteratorFromQuery(q)
	defer results.Release()

	config := csv.DefaultEncoderConfig()
	encoder := csv.NewMultiResultEncoder(config)
	b := bytes.Buffer{}
	_, err = encoder.Encode(&b, results)
	if err != nil {
		return "", err
	}
	results.Release()

	return string(b.Bytes()), results.Err()
}

func injectDependencies(ctx context.Context) context.Context {
	_deps := flux.NewDefaultDependencies()

	deps := Dependencies{
		Deps: _deps,

		influxdb: influxdb.Dependency{
			Provider: &HttpProvider{&influxdb.HttpProvider{DefaultConfig: influxdb.Config{}}},
		},

		bigtable: bigtable.Dependency{
			Provider: bigtable.DefaultProvider{},
		},

		mqtt: mqtt.Dependency{
			Dialer: mqtt.DefaultDialer{},
		},
	}
	ctx, span := dependency.Inject(ctx, deps)
	defer span.Finish()
	return ctx
}

type HttpProvider struct {
	*influxdb.HttpProvider
}

func (h *HttpProvider) ReaderFor(
	ctx context.Context, conf influxdb.Config, bounds flux.Bounds,
	predicateSet influxdb.PredicateSet) (influxdb.Reader, error) {
	if conf.Host == "" {
		return nil, fmt.Errorf("host is required")
	}
	return h.HttpProvider.ReaderFor(ctx, conf, bounds, predicateSet)
}

func (h *HttpProvider) SeriesCardinalityReaderFor(ctx context.Context,
	conf influxdb.Config, bounds flux.Bounds, predicateSet influxdb.PredicateSet) (influxdb.Reader, error) {
	if conf.Host == "" {
		return nil, fmt.Errorf("host is required")
	}
	return h.HttpProvider.SeriesCardinalityReaderFor(ctx, conf, bounds, predicateSet)
}

func (h *HttpProvider) WriterFor(ctx context.Context, conf influxdb.Config) (influxdb.Writer, error) {
	if conf.Host == "" {
		return nil, fmt.Errorf("host is required")
	}
	return h.HttpProvider.WriterFor(ctx, conf)
}

type Dependencies struct {
	flux.Deps
	influxdb influxdb.Dependency
	bigtable bigtable.Dependency
	mqtt     mqtt.Dependency
}

func (d Dependencies) Inject(ctx context.Context) context.Context {
	ctx = d.Deps.Inject(ctx)
	ctx = d.influxdb.Inject(ctx)
	ctx = d.bigtable.Inject(ctx)
	return d.mqtt.Inject(ctx)
}
