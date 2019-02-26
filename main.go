package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("/dist", "./dist")
	r.StaticFile("/", "./index.html")
	r.Run()
}
