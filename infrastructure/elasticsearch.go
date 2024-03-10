package elasticsearch

import (
	"fmt"

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
