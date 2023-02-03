package stats

import (
	"encoding/json"
	"maply/cache/managers/stats"
	"maply/models"
	friendsDBManager "maply/repository/managers/friends"
	"maply/ws"
	"time"
)

func UpdateStats(userId string, s *models.Stats) error {
	var now = time.Now().Unix()
	s.LastUpdate = int(now)

	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return stats.UpdateStats(userId, b)
}

func GetStats(userId string) {
	friends, err := friendsDBManager.GetFriendsId(userId)
	if err != nil || len(friends) == 0 {
		return
	}

	var friendsStats = make(map[string]*models.Stats)
	s, err := stats.GetFriendsStats(friends)
	for i := range s {
		friendId := friends[i]
		m := &models.Stats{}

		if s[i] != nil {
			err := json.Unmarshal([]byte(s[i].(string)), m)
			if err != nil {
				return
			}

			// Update online status
			if ws.GetClientConnection(friendId) != nil {
				m.IsOnline = true
			}
			friendsStats[friendId] = m
		}
	}

	ws.NewEvent(userId, ws.FriendsStats, friendsStats)
}
