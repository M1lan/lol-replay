package history

import (
	"encoding/json"
	"github.com/revel/revel"
	"io/ioutil"
	"time"
)

type Game struct {
	GameId        string
	Region        string
	EncryptionKey string
	Time          int64
}

func GameList() []Game {
	data, err := ioutil.ReadFile(revel.BasePath + "/history.json")
	if err != nil {
		return []Game{}
	}

	result := []Game{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		revel.ERROR.Println("Error reading games history")
		return []Game{}
	}

	return result
}

func StoreGame(region, gameId, encryptionKey string) {
	games := GameList()

	for _, value := range games {
		if region == value.Region && gameId == value.GameId {
			return
		}
	}

	game := Game{
		GameId:        gameId,
		Region:        region,
		EncryptionKey: encryptionKey,
		Time:          time.Now().Unix(),
	}

	newGames := []Game{game}

	games = append(newGames, games...)

	data, err := json.Marshal(games)
	if err != nil {
		revel.ERROR.Println("Error encoding games")
		return
	}

	err = ioutil.WriteFile(revel.BasePath+"/history.json", data, 0644)
	if err != nil {
		revel.ERROR.Println("Error saving game history!")
		return
	}
}
