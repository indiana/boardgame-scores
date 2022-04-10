package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indiana/boardgame-scores/models"
)

func getGames(c *gin.Context) {
	games, err := models.GetGames(10)
	checkErr(err)

	if games == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No records found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, games)
	}
}

func getGameById(c *gin.Context) {

	id := c.Param("id")

	game, err := models.GetGameById(id)
	checkErr(err)
	if game.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, game)
	}
}

func createGame(c *gin.Context) {

	var json models.Boardgame

	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.AddGame(json)

	if success {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func updateGame(c *gin.Context) {

	var json models.Boardgame

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.UpdateGame(json, gameId)

	if success {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func deleteGame(c *gin.Context) {

	gameId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.DeleteGame(gameId)

	if success {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func options(c *gin.Context) {

	ourOptions := "HTTP/1.1 200 OK\n" +
		"Allow: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Origin: http://locahost:8080\n" +
		"Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Headers: Content-Type\n"

	c.String(200, ourOptions)
}

func main() {
	err := models.ConnectDatabase()
	checkErr(err)
	router := gin.Default()
	router.GET("/game", getGames)
	router.GET("/game/:id", getGameById)
	router.POST("/game", createGame)
	router.PUT("/game/:id", updateGame)
	router.DELETE("/game/:id", deleteGame)
	router.OPTIONS("/game", options)
	router.Run("localhost:8080")
}
