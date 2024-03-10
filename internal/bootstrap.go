package internal

import (
	"log"

	elasticsearch "github.com/adrichard/siderproject/infrastructure"
	"github.com/adrichard/siderproject/internal/domain"
	"github.com/adrichard/siderproject/internal/ui"
	esClient "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

type Bootstrap struct {
	Config     esClient.Config
	EsClientV8 *elasticsearch.EsClient
	Router     *gin.Engine
}

func InitBootStrap(host, username, password string) Bootstrap {
	// I would have make a config.go file to handle the configuration
	EsClientV8, err := elasticsearch.NewEsClient(host, username, password)
	if err != nil {
		log.Fatal("Could not create client elasticsearch ", err)
		panic(err)
	}
	app := Bootstrap{
		EsClientV8: EsClientV8,
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

	//elasticsearch.IndexDocument(b.EsClientV8.Client, documents)
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
