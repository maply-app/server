package friends

import (
	"maply/repository"
)

func AddFriend(userId, friendId string) error {
	query := "INSERT INTO user_friends (user_id, friend_id) VALUES (?, ?), (?, ?);"
	err := repository.DB.Exec(query, userId, friendId, friendId, userId).Error
	return err
}

func DeleteFriend(userId, friendId string) error {
	query := "DELETE FROM user_friends WHERE user_id = ? AND friend_id = ? OR user_id = ? AND friend_id = ?;"
	err := repository.DB.Exec(query, userId, friendId, friendId, userId).Error
	return err
}

func CheckFriendByID(userId, friendId string) bool {
	var count int64
	query := "SELECT count(*) FROM user_friends WHERE user_id = ? AND friend_id = ?;"
	repository.DB.Raw(query, userId, friendId).Scan(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}

func GetFriendsId(userId string) ([]string, error) {
	var friends []string
	query := "SELECT friend_id FROM user_friends WHERE user_id = ?;"
	err := repository.DB.Raw(query, userId).Find(&friends).Error
	return friends, err
}
