package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// IndexHandler Handles default route
func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func init() {
	fmt.Println("Base Routes Initialized")
}
