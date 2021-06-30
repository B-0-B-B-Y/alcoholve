package player

type Player struct {
	PlayerID string `json:"playerId" binding:"required"`
	GameID   string `json:"gameId"`
	Name     string `json:"name" binding:"required"`
	Score    int32  `json:"score"`
}
