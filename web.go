package main

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
}

func (h *handler) health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func engine(h *handler) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.GET("/health", h.health)

	return engine
}
