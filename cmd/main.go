package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/go-kit/kit/log"
	"github.com/nukr/street_name/pkg/handler"
	"github.com/nukr/street_name/pkg/parser"
	"github.com/nukr/street_name/pkg/types"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var streetNames *types.Address
	{
		streetNames = parser.LoadAndParse("Xml_10510.xml")
	}

	var r http.Handler
	{
		r = handler.NewRouter(streetNames)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, r)
	}()
	logger.Log("exit", <-errs)
}
