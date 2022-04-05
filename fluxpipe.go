package main

import (
	"bufio"
	"os"
	"context"
//	"io"
//	"log"
	"fmt"
	"strings"
	"bytes"
	"time"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/dependency"
	"github.com/influxdata/flux/execute/executetest"
	_ "github.com/influxdata/flux/fluxinit/static"
	"github.com/influxdata/flux/lang"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/csv"
	"github.com/influxdata/flux/runtime"
)

var APPNAME = "flux-pipe"

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

var validScript = `
import "csv"

data = "
#datatype,string,long,long,string
#group,false,false,false,true
#default,_result,,,
,result,table,value,tag
,,0,10,a
,,0,10,a
,,1,20,b
,,1,20,b
,,2,30,c
,,2,30,c
,,3,40,d
,,3,40,d
"

csv.from(csv: data)  |> filter(fn: (r) => r["value"] >= 20) |> yield(name: "res") `

func main() {


	scanner := bufio.NewScanner((os.Stdin))
	inputString := ""
	for scanner.Scan() {
	        inputString = inputString + "\n" + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// fmt.Println(inputString)

	q, close, err := runQuery(context.Background(), inputString)
		if err != nil {
			fmt.Println("unexpected error while creating query: %s", err)
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


