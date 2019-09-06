package main

import (
	"fmt"
	"net/http"
	"os"

	"creart_new/controllers"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func main() {
	router := gin.Default()

	// group authentication create account creator
	auth := router.Group("/auth")
	{
		auth.POST("/v1/create-account", controllers.CreateCreatorAccount)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	router.Run()
	appengine.Main()
}
