package routes

import (
	"github.com/B-0-B-B-Y/alcoholve/server/game"
	"github.com/B-0-B-B-Y/alcoholve/server/player"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NewGameJson struct {
	PlayerNames []string `json:"playerNames" binding:"required"`
	Threshold   int32    `json:"threshold"`
	Alcohol     string   `json:"alcohol"`
}

// NewGame : Create a new game instance
func NewGame(c *gin.Context) {
	var inputData NewGameJson
	var gameData game.Game
	var playerList []player.Player

	c.BindJSON(&inputData)
	gameData = game.NewDefaultGame()
	playerCount := len(inputData.PlayerNames)

	for i := 0; i < playerCount; i++ {
		var playerData player.Player

		playerData.PlayerID = uuid.NewString()
		playerData.GameID = gameData.GameID
		playerData.Name = inputData.PlayerNames[i]
		playerData.Score = 0

		playerList = append(playerList, playerData)
	}
	gameData.PlayerList = playerList
	gameData.PlayerCount = int8(playerCount)

	if inputData.Alcohol != "" {
		gameData.Alcohol = inputData.Alcohol
	}

	if inputData.Threshold != 0 {
		gameData.Threshold = inputData.Threshold
	}

	c.JSON(200, gin.H{
		"game": gameData,
	})
}
