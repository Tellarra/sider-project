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

func NewEsClient(host, username, password string) (*EsClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{host},
		Username:  username,
		Password:  password,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &EsClient{Client: es}, nil
}

func CreateIndex(client *EsClient, indexName string) error {
	res, err := client.Client.Indices.Create(indexName)
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
