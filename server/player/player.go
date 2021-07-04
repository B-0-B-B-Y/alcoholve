package player

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Player : Player data object
type Player struct {
	PlayerID string `json:"playerId" binding:"required"`
	GameID   string `json:"gameId"`
	Name     string `json:"name" binding:"required"`
	Score    int32  `json:"score"`
}

// PlayerMove : An object representing a player's answer
type PlayerMove struct {
	PlayerID string `json:"playerId" binding:"required"`
	GameID   string `json:"gameId" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// PlayerWebsocketHandler : Create a websocket connection between the client (player) and the game server
func PlayerWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
