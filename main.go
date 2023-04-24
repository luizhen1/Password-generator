package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()

	router.GET("/password", func(c *gin.Context) {
		length := 12
		password := generatePassword(length)
		c.JSON(http.StatusOK, gin.H{"password": password})
	})

	router.Run(":8080")
}

func generatePassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
