package api

import "github.com/gin-gonic/gin"

func addRoutes(r *gin.Engine, store *Store) {
	api := r.Group("/api")
	api.GET("/todos", func(ctx *gin.Context) {})
	api.POST("/todos", func(ctx *gin.Context) {})
	api.DELETE("/todos/:id", func(ctx *gin.Context) {})
	api.PUT("/todos/:id", func(ctx *gin.Context) {})
}
