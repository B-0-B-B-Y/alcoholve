package redisClient

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/B-0-B-B-Y/alcoholve/server/game"
	"github.com/B-0-B-B-Y/alcoholve/server/player"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

// New : Create a new redis client
func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

// AddPlayerToGame : Add a new player to a game object in redis using GameID
func AddPlayerToGame(gameId string, playerName string) error {
	rdb := New()
	var gameData game.Game
	var newPlayer player.Player

	gameObject, err := rdb.Get(context.Background(), gameId).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(gameObject), &gameData)
	if err != nil {
		return err
	}

	for i := 0; i < len(gameData.PlayerList); i++ {
		if gameData.PlayerList[i].Name == playerName {
			return errors.New("a player with this name already exists, cannot join the game")
		}
	}

	newPlayer.GameID = gameId
	newPlayer.PlayerID = uuid.NewString()
	newPlayer.Name = playerName
	newPlayer.Score = 0

	gameData.PlayerList = append(gameData.PlayerList, newPlayer)
	gameData.PlayerCount += 1

	gameDataJSON, err := json.Marshal(gameData)
	if err != nil {
		return err
	}

	err = rdb.Set(context.Background(), gameData.GameID, gameDataJSON, 8.64e+13).Err()
	if err != nil {
		return err
	}

	return nil
}
