package main

import (
	"fmt"
	"time"

	"github.com/chadwwatson/sales-go/web/routes"
	"github.com/chadwwatson/sales-go/web/routes/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("web/views/**")
	r.Static("/assets", "web/assets")

	auth.Register(r)
	authorized := r.Group("/")
	authorized.Use(logMe())
	{
		authorized.GET("/", routes.IndexHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Run(":5000")
}

func logMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		fmt.Println("routes")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		fmt.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		fmt.Println(status)
	}
}
