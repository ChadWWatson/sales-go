package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Register initializes routes for authentication
func Register(router *gin.Engine) {
	fmt.Println("Registering Auth Routes")
	router.GET("/login", loginHandler)
}

func loginHandler(c *gin.Context) {
	c.HTML(200, "auth-login.tmpl", nil)
}
