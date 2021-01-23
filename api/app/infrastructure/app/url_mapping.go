package app

import (
	"fmt"
	"github.com/TestBackendMeli/api/app/infrastructure/config"
	"github.com/TestBackendMeli/api/app/infrastructure/container"
	"github.com/gin-gonic/gin"
)

func mapUrls(router *gin.Engine) {
	prefixScope := config.GetRoutePrefix()
	router.GET("/ping")
	prefix := fmt.Sprintf("%s/api", prefixScope)

	baseUrl := router.Group(prefix)
	content := baseUrl.Group("/meli")

	//content.GET("/", container.GetContentController().MakeGetContents)
	//content.GET("/all", container.GetContentController().MakeGetAllContent)
	content.POST("/topsecret/", container.GetQuasarController().PostTopSecret)
}
