package app

import (
	"fmt"
	"github.com/jhavierc/TestBackendMeli/api/app/infrastructure/config"
	"github.com/jhavierc/TestBackendMeli/api/app/infrastructure/container"
	"github.com/gin-gonic/gin"
)

func mapUrls(router *gin.Engine) {
	prefixScope := config.GetRoutePrefix()
	router.GET("/ping")
	prefix := fmt.Sprintf("%s/api", prefixScope)

	baseUrl := router.Group(prefix)
	content := baseUrl.Group("/meli")

	content.POST("/topsecret/", container.GetQuasarController().PostTopSecret)
	content.POST("/topsecret_split/:satellite_name", container.GetQuasarController().PostTopSecretSplit)
}
