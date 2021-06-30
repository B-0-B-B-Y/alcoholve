package routes

import (
	"github.com/B-0-B-B-Y/alcoholve/server/game"
	"github.com/B-0-B-B-Y/alcoholve/server/player"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewGame : Create a new game instance
func NewGame(c *gin.Context) {
	game := game.NewDefaultGame()
	p1 := player.Player{
		PlayerID: uuid.NewString(),
		GameID:   game.GameID,
		Name:     "Bob",
		Score:    0,
	}
	game.PlayerList = []player.Player{
		p1,
	}

	c.JSON(200, gin.H{
		"game": game,
	})
}
