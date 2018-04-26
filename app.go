package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/google/gops/agent"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	seapi "github.com/tokopedia/gosample/api"
	"github.com/tokopedia/gosample/hello"
	"github.com/tokopedia/gosample/setrg"
	"github.com/tokopedia/logging/tracer"
	"gopkg.in/tokopedia/grace.v1"
	"gopkg.in/tokopedia/logging.v1"
)

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	if err := agent.Listen(&agent.Options{}); err != nil {
		log.Fatal(err)
	}

	hwm := hello.NewHelloWorldModule()
	stm := setrg.NewSetrgModule()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/hello", hwm.SayHelloWorld)

	http.HandleFunc("/manusia-ganjil", stm.ManusiaGanjil)

	http.HandleFunc("/programmer-muda", stm.ProgrammerMuda)

	http.HandleFunc("/api/post", seapi.HandlePost)

	go logging.StatsLog()

	tracer.Init(&tracer.Config{Port: 8700, Enabled: true})

	log.Fatal(grace.Serve(":9000", nil))
}
