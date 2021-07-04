package game

import (
	"github.com/B-0-B-B-Y/alcoholve/server/player"
	"github.com/google/uuid"
)

// Game : Game state tracking object
type Game struct {
	GameID       string          `json:"gameId"`
	OpenTDBToken string          `json:"token"`
	Questions    []GameCard      `json:"questions"`
	PlayerCount  int8            `json:"playerCount"`
	Round        int8            `json:"round"`
	Threshold    int32           `json:"threshold"`
	Alcohol      string          `json:"alcohol"`
	PlayerList   []player.Player `json:"playerList"`
}

// DefaultNewGame : Returns a game object with default values set for a new game
func NewDefaultGame() Game {
	return Game{
		GameID:       uuid.NewString(),
		OpenTDBToken: RequestToken().Token,
		PlayerCount:  3,
		Round:        1,
		Threshold:    500,
		Alcohol:      "Beer/Cider",
	}
}
