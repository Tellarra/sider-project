package elasticsearch

import (
	"bytes"
	"fmt"
	"log"

	"github.com/adrichard/siderproject/internal/model"
	"github.com/elastic/go-elasticsearch/v8"
)

type EsClient struct {
	Client *elasticsearch.Client
}

func CreateIndex(client *elasticsearch.Client, indexName string) error {
	res, err := client.Indices.Create(indexName)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error creating index: %s", res.String())
	}
	return nil
}

func IndexDocument(client *elasticsearch.Client, docs []model.DocumentToIndex) error {
	var task []byte
	for _, doc := range docs {
		// convert to JSON
		switch doc.Name {
		case "task":
			task = doc.Data
			/* if err != nil {
				return fmt.Errorf("error marshalling task: %w", err)
			} */
		}
	}

	res, err := client.Index(
		"tasks",               // Index name
		bytes.NewReader(task), // Document body
	)

	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	// Handle successful or failed indexing
	if res.StatusCode != 201 {
		log.Printf("Error indexing document: %s", res.String())
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	fmt.Println("Document indexed successfully:", res.Body)
	return nil
}
