package game

import (
	"github.com/B-0-B-B-Y/alcoholve/server/player"
	"github.com/google/uuid"
)

// Game : Game state tracking object
type Game struct {
	GameID      string
	PlayerCount int8
	Round       int8
	Threshold   int32
	Alcohol     string
	PlayerList  []player.Player
}

// DefaultNewGame : Returns a game object with default values set for a new game
func NewDefaultGame() Game {
	return Game{
		GameID:      uuid.NewString(),
		PlayerCount: 3,
		Round:       1,
		Threshold:   500,
		Alcohol:     "Beer/Cider",
	}
}
