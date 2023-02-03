package stats

import (
	"maply/cache"
	"maply/config"
)

func UpdateStats(userId string, stats []byte) error {
	return cache.Redis.Set(userId, stats, config.C.Stats.TTL).Err()
}

func GetStats(userId string) (interface{}, error) {
	s, err := cache.Redis.Get(userId).Result()
	if err != nil {
		return s, err
	}
	return s, err
}

func GetFriendsStats(friends []string) ([]interface{}, error) {
	s, err := cache.Redis.MGet(friends...).Result()
	if err != nil {
		return s, err
	}
	return s, err
}
