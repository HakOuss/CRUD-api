package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/person/CRUD-api/models"
	"github.com/person/CRUD-api/services"
)

func AddPerson(services *services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json models.Person
		if ctx.BindJSON(&json) != nil {
			return
		}
		if json.Name == "" || json.LastName == "" {
			models.MissingError(ctx)
			return
		}
		services.PersonService.Create(&json)
		ctx.JSON(http.StatusOK, gin.H{"status": "Person added successfully."})
	}
}

func ListPerson(services *services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		list := services.PersonService.List()
		if len(list) == 0 {
			models.EmptyList(ctx)
			return
		}
		ctx.JSON(http.StatusOK, list)
	}
}
func UpdatePerson(services *services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("index")
		index, err := strconv.Atoi(param)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "Invalid index please put an integer!"})
			return
		}
		var newP models.Person
		p := services.PersonService.List()[index]
		if p == nil {
			models.NotFound(ctx)
			return
		}
		if ctx.BindJSON(&newP) != nil {
			return
		}
		if newP.Name == "" && newP.LastName == "" {
			models.MissingError(ctx)
			return
		}
		if newP.Name == "" {
			newP.Name = p.Name
		}
		if newP.LastName == "" {
			newP.LastName = p.LastName
		}
		services.PersonService.Update(index, &newP)
		ctx.JSON(http.StatusOK, gin.H{"status": "Person updated successfully"})
	}
}
func DeletePerson(services *services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("index")
		index, err := strconv.Atoi(param)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "Invalid index please put an integer!"})
			return
		}
		p := services.PersonService.List()[index]
		if p == nil {
			models.NotFound(ctx)
			return
		}
		services.PersonService.Delete(index)
		ctx.JSON(http.StatusOK, gin.H{"status": "Person deleted successfully"})
	}
}
