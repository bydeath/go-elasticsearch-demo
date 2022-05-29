package main

import (
	"es-demo/util"
	"io/ioutil"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Println(config)
	cert, _ := ioutil.ReadFile(config.ESCacertpath)
	cfg := elasticsearch.Config{
		Addresses: config.ESNodes,
		Username:  config.ESUsername,
		Password:  config.ESPassword,
		CACert:    cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	log.Println(es.Info())

	res, err := es.Index(
		"test",                                  // Index name
		strings.NewReader(`{"title" : "Test"}`), // Document body
		es.Index.WithDocumentID("1"),            // Document ID
		es.Index.WithRefresh("true"),            // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
}
