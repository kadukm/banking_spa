package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/handling"
)

func main() {
	useCommonEngine := true
	if useCommonEngine {
		go runCommonEngine()
	}
	runAPIEngine()
}

func runCommonEngine() {
	commonEngine := gin.Default()
	buildCommonRoutes(commonEngine)
	commonEngine.Run(":8080")
}

func runAPIEngine() {
	apiEngine := gin.Default()
	buildAPIRoutes(apiEngine)
	apiEngine.Run(":3000")
}

func buildCommonRoutes(engine *gin.Engine) {
	indexHandler := func(c *gin.Context) {
		c.File("./index.html")
	}
	engine.Static("/assets", "./assets")
	engine.StaticFile("/admin-panel", "./index.html")
	engine.StaticFile("/", "./index.html")
	engine.GET("/companies/:companyID", indexHandler)
	engine.HEAD("/companies/:companyID", indexHandler)
}

func buildAPIRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
	api.Use(corsMiddleware())
	{
		payments := api.Group("/payments")
		{
			payments.POST("/from_card", handling.PostPaymentFromCard)
			payments.POST("/requests", handling.PostPaymentRequest)
			payments.POST("/via_bank", handling.PostPaymentViaBank)
			payments.PATCH("/from_card/:paymentID", handling.PatchPaymentFromCard)
		}
		companies := api.Group("/companies")
		{
			companies.GET("/:companyID", handling.GetCompany)
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	}
}
