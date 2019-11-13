package controllers

import (
	"net/http"
	"user_services/models"

	"github.com/gin-gonic/gin"
)

// Api for auth new user registration
var CreateCreatorAccount = func(c *gin.Context) {
	json := &models.AccountData{}

	if err := c.ShouldBindJSON(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	json.Source = "Create from Apps"

	response := json.CreateCreator()
	c.JSON(200, response)
}

var Test = func(c *gin.Context) {
	json := &models.AccountData{}

	if err := c.ShouldBindJSON(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	response := json.Data()
	c.JSON(200, response)
}

var AuthLogin = func(c *gin.Context) {
	json := &models.AccountData{}

	if err := c.ShouldBindJSON(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	response := gin.H{
		"statusCode": "00",
		"username":   json.Username,
	}
	c.JSON(200, response)
}
