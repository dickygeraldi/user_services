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
	// group authentication create account creator
	// router.HandleFunc("/v1/create-account", controllers.CreateCreatorAccount).Methods("POST") // Create Creator Account
	router.POST("/v1/create-account", controllers.CreateCreatorAccount)

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
