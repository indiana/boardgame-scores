package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bg "github.com/indiana/boardgame-scores/boardgame"
)

var games = []bg.Boardgame{
	{ID: 1, Name: "Zaginiona wyspa Arnak", Description: "Zaginiona wyspa Arnak"},
	{ID: 2, Name: "Everdell", Description: "Everdell"},
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func createGame(c *gin.Context) {
	var newGame bg.Boardgame

	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.POST("/games", createGame)
	router.Run("localhost:8080")
}
