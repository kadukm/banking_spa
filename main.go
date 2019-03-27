package main

import (
	"flag"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/handling"
	"github.com/kadukm/banking_spa/server/utils"
)

func main() {
	dev := flag.Bool("dev", false, "a bool")
	flag.Parse()

	if *dev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowHeaders = []string{"Content-Type", utils.CSRFTokenName}
	config.AllowCredentials = true
	apiEngine.Use(cors.New(config))
	buildAPIRoutes(apiEngine)
	apiEngine.Run(":3000")
}

func buildCommonRoutes(engine *gin.Engine) {
	indexHandler := func(c *gin.Context) {
		c.File("./index.html")
	}
	engine.Static("/assets", "./assets")
	engine.GET("/admin-panel", utils.CSRFGeneration, indexHandler)
	engine.GET("/", utils.CSRFGeneration, indexHandler)
	engine.GET("/companies/:companyID", utils.CSRFGeneration, indexHandler)
}

func buildAPIRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		payments := api.Group("/payments")
		{
			payments.GET("/from_card", handling.GetPaymentsFromCard)
			payments.GET("/from_card/sort", handling.GetPaymentsFromCardSorted)
			payments.GET("/requests", handling.GetPaymentRequests)
			payments.GET("/requests/sort", handling.GetPaymentRequestsSorted)

			payments.PATCH("/from_card/:paymentID", utils.CheckCSRFToken, handling.PatchPaymentFromCard)

			payments.POST("/from_card", utils.CheckCSRFToken, handling.PostPaymentFromCard)
			payments.POST("/requests", utils.CheckCSRFToken, handling.PostPaymentRequest)
			payments.GET("/via_bank", handling.GetPaymentViaBank)
		}
		companies := api.Group("/companies")
		{
			companies.GET("/:companyID", handling.GetCompany)
			companies.GET("/:companyID/products", handling.GetProducts)
		}
	}
}
