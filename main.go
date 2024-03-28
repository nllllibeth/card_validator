package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"card_validator/validator"
)

func main () {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/validate", func(c * gin.Context) {
		cardNumber := c.PostForm("card_number")
		valid := validator.IsValidCreditCardNumber(cardNumber)

		var message string
		if valid {
			message = "Valid credit card number!"
		} else {
			message = "Invalid credit card number!"
		}

		c.HTML(http.StatusOK, "result.html", gin.H{"message": message})
	})
	r.Run(":8080")
}