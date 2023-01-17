package stats

import (
	"encoding/json"
	"maply/cache/manager/stats"
	"maply/models"
	"maply/repository/managers"
	"maply/ws"
)

func UpdateStats(userId string, s *models.Stats) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return stats.UpdateStats(userId, b)
}

func GetStats(userId string) {
	friends, err := managers.GetFriendsId(userId)
	if err != nil || len(friends) == 0 {
		return
	}

	var friendsStats = make(map[string]*models.Stats)
	s, err := stats.GetFriendsStats(friends)
	for i := range s {
		friendsId := friends[i]
		m := &models.Stats{}

		if s[i] != nil {
			json.Unmarshal([]byte(s[i].(string)), m)

			// Update online status
			if ws.GetClientConnection(friendsId) != nil {
				m.IsOnline = true
			}
			friendsStats[friendsId] = m
		}
	}

	// Send socket event
	ws.NewEvent(userId, ws.FriendsStats, friendsStats)
}
