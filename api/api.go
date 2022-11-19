package api

import (
	"github.com/gin-gonic/gin"

	"github.com/digvijay17july/SqlDbToCSV/ui"
)

func Start() {
	var store = &Store{}

	router := gin.Default()

	addRoutes(router, store)
	ui.AddRoutes(router)

	router.Run(":4000")
}
