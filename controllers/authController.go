package controllers

import (
	"creart_new/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Api for auth new user registration
var CreateCreatorAccount = func(c *gin.Context) {
	json := &models.AccountData{}

	if err := c.ShouldBindJSON(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	response := json.CreateCreator()
	c.JSON(200, response)
}
