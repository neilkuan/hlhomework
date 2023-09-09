package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/orian/counters"
)

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(html)
	r.StaticFile("/favicon.ico", "./favicon.ico")
	cb := counters.NewCounterBox()
	pageViewCounter := cb.GetCounter("pageView")

	r.GET("/", func(c *gin.Context) {
		pageViewCounter.Increment()
		c.HTML(http.StatusOK, "index", gin.H{
			"count": pageViewCounter.Value(),
		})
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
