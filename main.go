package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/handling"
)

func main() {
	engine := gin.Default()
	buildCommonRoutes(engine)
	buildAPIRoutes(engine)
	engine.Run()
}

func buildCommonRoutes(engine *gin.Engine) {
	indexHandler := func(c *gin.Context) {
		c.File("./index.html")
	}
	engine.Static("/dist", "./dist")
	engine.StaticFile("/admin-panel", "./index.html")
	engine.StaticFile("/", "./index.html")
	engine.GET("/companies/:companyID", indexHandler)
	engine.HEAD("/companies/:companyID", indexHandler)
}

func buildAPIRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
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
