package ui

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	HostAddress string
	HTTPPort    int
}

func InitRouterConfig(
	routerConfig RouterConfig,
	es *elasticsearch.Client,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", HealthCheck())
	r.GET("/tasks",
		func(ctx *gin.Context) {
			GetTasks(ctx, es)
		},
	)
	r.GET("/feed",
		func(ctx *gin.Context) {
			Feed(ctx, es)
		})

	return r
}
