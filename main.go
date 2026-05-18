package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"Repositories/E-COM/controllers"
	"Repositories/E-COM/databases"
	"Repositories/E-COM/middleware"
	"Repositories/E-COM/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(databases.ProductData(databases.Cilent, "Products"),
		databases.UserData(databases.Cilent, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
