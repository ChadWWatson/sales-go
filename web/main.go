package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("web/views/*")
    r.Static("/assets", "web/assets")
    r.GET("/",indexHandler)
    r.Run(":5000") // listen and server on 0.0.0.0:8080
}

func indexHandler(c *gin.Context) {
    c.HTML(200, "index.html", nil)
}
