package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/olivere/elastic"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"facility":{
			"properties":{
				"type":{
					"type":"text"
				},
				"nature":{
					"type":"text"
				},
				"city":{
					"type":"text"
				},
				"commissioning_year":{
					"type":"text"
				},
				"name":{
					"type":"text"
				},
				"address":{
					"type":"text"
				},
				"nb_facilities":{
					"type":"integer"
				},
				"facility_id":{
					"type":"integer"
				},
				"zip_code":{
					"type":"integer"
				}
			}
		}
	}
}`

func main() {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("facilities").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("facilities").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()["q"]
		mq := elastic.NewMatchQuery("name", q[0])
		searchResults, err := client.Search().Index("facilities").Query(mq).Do(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if searchResults.Hits.TotalHits <= 0 {
			fmt.Fprint(w, "no result found")
			return
		}

		fmt.Fprint(w, "%q", searchResults.Hits)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
