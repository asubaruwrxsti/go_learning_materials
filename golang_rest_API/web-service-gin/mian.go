package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Car struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var cars = []Car{
	{ID: "1", Name: "BMW"},
	{ID: "2", Name: "Mercedes"},
	{ID: "3", Name: "Audi"},
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.GET("/cars/:id", getCarByID)
	router.POST("/cars", postCar)

	router.Run("localhost:8080")
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func postCar(c *gin.Context) {
	var newCar Car

	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func getCarByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range cars {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}