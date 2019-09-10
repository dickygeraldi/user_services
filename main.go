package main

import (
	"fmt"
	"net/http"
	"os"

	"user_services/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := mux.NewRouter()
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/v1/create-account", controllers.CreateCreatorAccount)
		auth.POST("/v1/coba", controllers.Test)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	// router.Use(authentication.JwtAuthentication)
}
