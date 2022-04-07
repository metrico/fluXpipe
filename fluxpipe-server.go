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
	"encoding/json"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/dependency"
	"github.com/influxdata/flux/execute/executetest"
	_ "github.com/influxdata/flux/fluxinit/static"
	"github.com/influxdata/flux/lang"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/csv"
	"github.com/influxdata/flux/runtime"

	"io/ioutil"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func postQuery(c echo.Context) error {

	c.Response().Header().Set(echo.HeaderContentType, "text/csv; charset=utf-8")
	c.Response().Header().Set("x-fluxpipe-cloud", "qxip")

	content := c.Request().Header.Get("Content-Type")
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	if strings.Contains(string(s), "buckets()") {
		// fake bucket to make grafana happy
		buckets := "#datatype,string,string,string,string,string,string,long\n"+
			   "#default,_result,,,,,,\n" +
			   ",result,table,name,id,organizationID,retentionPolicy,retentionPeriod\n"+
			   ",_result,0,_fluxpipe,aa9f5aa08895152b,03dbe8db13d17000,,604800000000000\n"+
			   "\n"
		return c.String(http.StatusOK, buckets)
	}

	if strings.Contains(content, "json") {
		json_map := make(map[string]interface{})
		err := json.Unmarshal(s, &json_map)
		if err != nil {
			return err
		} else {
			q := json_map["query"]
			query := fmt.Sprintf("%v", q)
			res := exec(query)
			return c.String(http.StatusOK, res)
		}

	} else {

		res := exec(string(s))
		return c.String(http.StatusOK, res)
	}
}

func exec(inputString string) string {

	q, close, err := runQuery(context.Background(), inputString)
	if err != nil {
		fmt.Println("unexpected error while creating query: %s", err)
		fmt.Println("%s", inputString)
		return string(fmt.Sprintf("%v", err))
	}
	defer close()

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
		return string(fmt.Sprintf("%v", q.Err()))

	} else {
		return buf.String()
	}
}

func main() {

	url  := flag.String("url", "", "ClickHouse MYSQL API URL")
	port := flag.String("port", "8086", "API port")
	stdin  := flag.Bool("stdin", false, "STDIN mode")
	cors  := flag.Bool("cors", true, "API cors mode")
	flag.Parse()

	scanner := bufio.NewScanner((os.Stdin))
	inputString := ""

	if *stdin == true {

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

		buf := exec(inputString)
		fmt.Println(strings.Replace(buf, "\r\n", "\n", -1))

	} else {

		e := echo.New()
		e.HideBanner = true
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		if *cors == true {
			e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
			}))
		}

		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "F-L-U-X-P-I-P-E")
		})
		e.GET("/ping", func(c echo.Context) error {
			return c.String(204, "OK")
		})

		e.POST("/api/v2/query", postQuery)
		e.POST("/query", postQuery)

		fmt.Println("|> FluxPIPE |>")
		e.Logger.Fatal(e.Start(":"+*port))
	}
}

