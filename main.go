package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Static("/dist", "./dist")
	r.StaticFile("/admin-panel", "./index.html")
	r.StaticFile("/", "./index.html")
	r.GET("/companies/:companyId", indexHandler)
	r.HEAD("/companies/:companyId", indexHandler)

	r.Run()
}

func indexHandler(c *gin.Context) {
	c.File("./index.html")
}
