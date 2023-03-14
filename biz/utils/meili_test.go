package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/meilisearch/meilisearch-go"
)

var C *meilisearch.Client

func init() {
	C = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "masterKey",
	})
}
func TestCreateDoc(t *testing.T) {

	index := C.Index("movies")
	documents := []map[string]interface{}{
		{"id": 1, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 2, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 3, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 4, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 5, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 6, "title": "Philadelphia", "genres": []string{"Drama"}},
	}
	task, err := index.AddDocuments(documents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(task.TaskUID)
}

func TestSearch(t *testing.T) {
	searchRes, err := C.Index("movies").Search("m",
		&meilisearch.SearchRequest{
			Limit: 10,
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hits := searchRes.Hits
	fmt.Println(hits)
	hitsJson, _ := json.Marshal(hits)
	fmt.Println(string(hitsJson))
}
