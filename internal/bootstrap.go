package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/adrichard/siderproject/internal/domain"
	"github.com/adrichard/siderproject/internal/model"
	"github.com/adrichard/siderproject/internal/ui"
	"github.com/elastic/go-elasticsearch/esapi"
	esClient "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

type Bootstrap struct {
	Config     esClient.Config
	EsClientV8 *esClient.Client
	Router     *gin.Engine
}

func InitBootStrap() Bootstrap {
	esClientV8, err := esClient.NewDefaultClient()
	if err != nil {
		log.Fatal("Could not create elasticsearch client ", err)
	}
	app := Bootstrap{
		Config: esClient.Config{
			Addresses: []string{
				"http://localhost:9200",
			},
		},
		EsClientV8: esClientV8,
		Router:     gin.Default(),
	}
	return app
}

func (b Bootstrap) Run() {
	b.CreateRouter()
	if b.Router == nil {
		log.Fatal("Router is not initialized")
	}

	if err := b.Router.Run(":8080"); err != nil {
		log.Fatal("Could not start the server ", err)
	}
}

func (b Bootstrap) CreateRouter() {
	documents, err := domain.GetAllFilesToRightTypes()
	if err != nil {
		log.Fatal("Could not get files ", err)
	}

	// Create index
	b.DatabaseOperations(documents)
	b.GetDatas()

	b.Router.GET("/health", b.HealthCheck)
	b.Router.GET("/tasks", func(ctx *gin.Context) {
		ui.GetTasks(ctx, documents)
	})
	b.Router.PATCH("/tasks/:id", func(ctx *gin.Context) {
		ui.UpdateAssigneeID(ctx, documents)
	})
}

func (b Bootstrap) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I'm alive",
	})
}

/* func (b Bootstrap) DatabaseOperations(docs []model.DocumentToIndex) {
	var buf bytes.Buffer
	for _, doc := range docs {
		if doc.Name != "shift" {
			continue
		}
		for _, datas := range doc.Data {
			meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "%s" } }%s`, doc.Name, "\n"))
			data, err := json.Marshal(datas)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, "\n"...) // Ajouter une nouvelle ligne à la fin de chaque document

			buf.Write(meta)
			buf.Write(data)
		}
	}

	res, err := b.EsClientV8.Bulk(bytes.NewReader(buf.Bytes()), b.EsClientV8.Bulk.WithIndex("shift"))
	if err != nil {
		log.Fatal(err)
	}
	if res.IsError() {
		log.Fatalf("bulk indexation failed: %s", res.String())
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Errorf("bulk indexation failed: %s", res.String())
	}

	log.Println("Bulk indexation succeeded")
} */

func (b Bootstrap) DatabaseOperations(docs []model.DocumentToIndex) {
	for _, doc := range docs {
		if doc.Name != "task" {
			continue
		}

		var tasks []model.Task
		err := json.Unmarshal(doc.Data, &tasks)
		if err != nil {
			log.Fatalf("Error unmarshalling task: %v", err)
		}

		for _, task := range tasks {
			// Convertit l'objet task en JSON pour l'indexation
			jsonData, err := json.Marshal(task)
			if err != nil {
				log.Fatalf("Error marshalling task: %v", err)
			}

			// Créez une requête d'indexation pour chaque document
			req := esapi.IndexRequest{
				Index:   "tasks",
				Body:    bytes.NewReader(jsonData),
				Refresh: "true",
			}

			// Exécutez la requête d'indexation
			res, err := req.Do(context.Background(), b.EsClientV8)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("Error indexing task ID %s: %s", task.ID, res.String())
			} else {
				log.Printf("Task ID %s indexed.", task.ID)
			}
		}
	}
}

func (b Bootstrap) GetDatas() {
	// Construisez votre requête de recherche. Ceci est un exemple simple de match_all.
	var buf strings.Builder
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Exécutez la recherche
	res, err := b.EsClientV8.Search(
		b.EsClientV8.Search.WithContext(context.Background()),
		b.EsClientV8.Search.WithIndex("task"),
		b.EsClientV8.Search.WithBody(bytes.NewBufferString(buf.String())),
		b.EsClientV8.Search.WithTrackTotalHits(true),
		b.EsClientV8.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Affichez l'erreur spécifique à Elasticsearch.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	result := model.Task{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(result)
}
