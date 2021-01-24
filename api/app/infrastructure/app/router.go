package app

import (
	"os"
	"github.com/jhavierc/TestBackendMeli/api/app/infrastructure/controller/middleware"
	"github.com/jhavierc/TestBackendMeli/goutils/logger"
	"github.com/jhavierc/TestBackendMeli/goutils/mlhandlers"
)

func StartApp() {
	router := mlhandlers.DefaultRouter()
	router.Use(middleware.ErrorHandler())
	mapUrls(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "3000"
	}

	if err := router.Run(port); err != nil {
		logger.Error("error running server", err)
	}
}
