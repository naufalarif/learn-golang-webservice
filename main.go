package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Game struct {
	ID 		string	`json:"id"`
	Title string 	`json:"title"`
	Developer	string	`json:"developer"`
}

var games = []Game {
	{ ID: "1", Title: "Honkai Impact 3", Developer: "Hoyoverse" },
	{ ID: "2", Title: "Genshin Impact 3", Developer: "Hoyoverse" },
	{ ID: "3", Title: "Dota 2", Developer: "Valve" },
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func postGame(c *gin.Context) {
	var newGame Game

	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

func getGameByID(c *gin.Context) {
	id := c.Param("id")

	for _, game := range games {
		if game.ID == id {
			c.IndentedJSON(http.StatusOK, game)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "game not found" })
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/games/:id", getGameByID)
	router.POST("/games/add", postGame)

	router.Run("localhost:8080")
}