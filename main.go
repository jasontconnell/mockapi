package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jasontconnell/mockapi/conf"
	"github.com/jasontconnell/mockapi/process"
)

func main() {
	c := flag.String("c", "config.json", "config filename")
	flag.Parse()

	cfg := conf.LoadConfig(*c)

	m := make(map[string]string)
	for _, mapping := range cfg.Mappings {
		m[mapping.Source] = mapping.Destination
	}

	s := process.NewServer(cfg.BasePath, m, cfg.ContentType)

	log.Fatal(http.ListenAndServe(cfg.Binding, s))
}
