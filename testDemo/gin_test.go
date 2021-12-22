package testDemo

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	r := gin.Default()

	r.Handle("GET", "/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9090")
}
