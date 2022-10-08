package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}
type PersonList map[int]*Person

func MissingError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": "Missing name or last name!"})
}
func EmptyList(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": "Person list is empty!"})
}
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"status": "Person doesn't exist!"})
}
