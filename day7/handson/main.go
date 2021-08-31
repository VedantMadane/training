package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("/src/gin-app/ls-1/02_static_routes_and_file_handlers/templates/**/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	admin := r.Group("/admin")
	admin.GET("/userlist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "/src/gin-app/ls-1/02_static_routes_and_file_handlers/public")

	r.Run(":3000")
}
