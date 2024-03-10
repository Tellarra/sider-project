package internal

import (
	"database/sql"
	"log"

	"github.com/adrichard/siderproject/internal/domain"
	"github.com/adrichard/siderproject/internal/ui"
	esClient "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

type Bootstrap struct {
	Config   esClient.Config
	DataBase *sql.DB
	Router   *gin.Engine
}

func InitBootStrap() Bootstrap {
	// Initialisation de SQLite
	db, err := sql.Open("sqlite3", "./data/database.sqlite")
	if err != nil {
		log.Fatal("Could not open SQLite database: ", err)
		panic(err) // Panic est généralement utilisé dans l'initialisation. Considérez une gestion d'erreur plus élaborée selon votre cas.
	}

	// Vous pouvez choisir de tester la connexion avec db.Ping() ici
	db.Ping()
	app := Bootstrap{
		DataBase: db,
		Router:   gin.Default(),
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
