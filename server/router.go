package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Nothing here")
	})

	r.POST("/getVulInfo", func(c *gin.Context) {
		s.getVulInfo(c)
	})

	return r
}
