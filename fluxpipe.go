package main

import (
	"bufio"
	"os"
	"context"
	"fmt"
	"strings"
	"bytes"
	"time"
	"flag"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/dependency"
	"github.com/influxdata/flux/execute/executetest"
	_ "github.com/influxdata/flux/fluxinit/static"
	"github.com/influxdata/flux/lang"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/csv"
	"github.com/influxdata/flux/runtime"
)

var APPNAME = "fluxpipe"

func runQuery(ctx context.Context, script string) (flux.Query, func(), error) {
	program, err := lang.Compile(script, runtime.Default, time.Unix(0, 0))
	if err != nil {
		return nil, nil, err
	}
	ctx, deps := dependency.Inject(ctx, executetest.NewTestExecuteDependencies())

	q, err := program.Start(ctx, memory.DefaultAllocator)
	if err != nil {
		deps.Finish()
		return nil, nil, err
	}
	return q, deps.Finish, nil
}

func main() {

	url := flag.String("url", "", "ClickHouse MySQL URL")
	flag.Parse()

	scanner := bufio.NewScanner((os.Stdin))
	inputString := ""

	for scanner.Scan() {
	        inputString = inputString + "\n" + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	if len(*url) > 1 {
		orig := `"clickhouse"`
		repl := fmt.Sprintf(`"mysql", dataSourceName: "%s"`, *url)
		inputString = strings.ReplaceAll(inputString, orig, repl)
	}

	q, close, err := runQuery(context.Background(), inputString)
	if err != nil {
		fmt.Println("unexpected error while creating query: %s", err)
		fmt.Println("%s", inputString)
	}
	defer close()

	results := flux.NewResultIteratorFromQuery(q)
	defer results.Release()

	buf := bytes.NewBuffer(nil)
	encoder := csv.NewMultiResultEncoder(csv.DefaultEncoderConfig())

	if _, err := encoder.Encode(buf, results); err != nil {
		panic(err)
	}

	// This substitution is done because the testable example's Output
	// section cannot contain carriage return while the csv encoder emits them
	fmt.Println(strings.Replace(buf.String(), "\r\n", "\n", -1))

	// release query resources
	q.Done()

	if q.Err() != nil {
		fmt.Println("unexpected error from query execution: %s", q.Err())
	}
}
