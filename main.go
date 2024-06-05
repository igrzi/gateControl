package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Any("/*path", gatewayHandler)

	router.Run(":8000")
}

func gatewayHandler(c *gin.Context) {
	path := c.Param("path")
	var target *url.URL
	var err error

	switch {
	case strings.HasPrefix(path, "/acesso"):
		target, err = url.Parse("http://localhost:8010")
	case strings.HasPrefix(path, "/usuario"):
		target, err = url.Parse("http://localhost:8020")
	case strings.HasPrefix(path, "/creditos"):
		target, err = url.Parse("http://localhost:8030")
	case strings.HasPrefix(path, "/vagas"):
		target, err = url.Parse("http://localhost:8040")
	case strings.HasPrefix(path, "/cancela"):
		target, err = url.Parse("http://localhost:8050")
	default:
		c.JSON(http.StatusNotFound, gin.H{"error": "Unknown path"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse target URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Writer, c.Request)
}
