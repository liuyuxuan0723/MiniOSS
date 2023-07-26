package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "hello world",
	})
}
