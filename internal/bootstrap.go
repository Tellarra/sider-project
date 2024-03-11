package internal

import (
	"log"

	"github.com/adrichard/siderproject/internal/ui"
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
	router := ui.InitRouterConfig(ui.RouterConfig{
		HostAddress: "localhost",
		HTTPPort:    8080,
	}, esClientV8)

	app := Bootstrap{
		Config:     esClient.Config{},
		EsClientV8: esClientV8,
		Router:     router,
	}
	return app
}

func (b Bootstrap) Run() {
	if b.Router == nil {
		log.Fatal("Router is not initialized")
	}

	if err := b.Router.Run(b.Config.Addresses...); err != nil {
		log.Fatal("Could not start the server ", err)
	}
}
