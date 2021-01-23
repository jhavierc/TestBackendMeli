package app

import (
	"os"
	"github.com/TestBackendMeli/api/app/infrastructure/controller/middleware"
	"github.com/TestBackendMeli/goutils/logger"
	"github.com/TestBackendMeli/goutils/mlhandlers"
)

func StartApp() {
	router := mlhandlers.DefaultRouter()
	router.Use(middleware.ErrorHandler())
	mapUrls(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "8085"
	}

	if err := router.Run(port); err != nil {
		logger.Error("error running server", err)
	}
}
