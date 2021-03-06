package routes

import (
	"context"
	"encoding/json"
	"log"

	"github.com/B-0-B-B-Y/alcoholve/server/game"
	"github.com/B-0-B-B-Y/alcoholve/server/player"
	redisClient "github.com/B-0-B-B-Y/alcoholve/server/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewGameJson : Input parameters when a client makes a new game request
type NewGameJson struct {
	PlayerNames        []string `json:"playerNames" binding:"required"`
	Threshold          int32    `json:"threshold"`
	Alcohol            string   `json:"alcohol"`
	QuestionAmount     string   `json:"questionAmount"`
	QuestionCategory   string   `json:"questionCategory"`
	QuestionDifficulty string   `json:"questionDifficulty"`
}

// NewGame : Create a new game instance
func NewGame(c *gin.Context) {
	var inputJSON NewGameJson
	var gameData game.Game
	var playerList []player.Player
	rdb := redisClient.New()

	// Read JSON data from request body and create initial game parameters
	c.BindJSON(&inputJSON)
	gameData = game.NewDefaultGame()
	playerCount := len(inputJSON.PlayerNames)

	// Create the player objects
	for i := 0; i < playerCount; i++ {
		var playerData player.Player

		playerData.PlayerID = uuid.NewString()
		playerData.GameID = gameData.GameID
		playerData.Name = inputJSON.PlayerNames[i]
		playerData.Score = 0

		playerList = append(playerList, playerData)
	}
	gameData.Questions = game.GetQuestions(
		gameData.OpenTDBToken,
		inputJSON.QuestionAmount,
		inputJSON.QuestionCategory,
		inputJSON.QuestionDifficulty,
	)
	gameData.PlayerList = playerList
	gameData.PlayerCount = int8(playerCount)

	if inputJSON.Alcohol != "" {
		gameData.Alcohol = inputJSON.Alcohol
	}

	if inputJSON.Threshold != 0 {
		gameData.Threshold = inputJSON.Threshold
	}

	// Store the newly created game and its state in redis
	gameDataJSON, err := json.Marshal(gameData)
	if err != nil {
		log.Panic("Failed to marshal game data.")
	}
	err = rdb.Set(context.Background(), gameData.GameID, gameDataJSON, 8.64e+13).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"game": gameData,
	})
}
