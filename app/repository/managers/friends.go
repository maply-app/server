package managers

import (
	"maply/models"
	"maply/repository"
)

// AddFriend ...
func AddFriend(userID, friendID string) error {
	query := "INSERT INTO user_friends (user_id, friend_id) VALUES (?, ?), (?, ?);"
	err := repository.DB.Exec(query, userID, friendID, friendID, userID).Error
	return err
}

// GetFriends ...
func GetFriends(userID string) ([]models.User, error) {
	var friends []models.User
	query := "SELECT * FROM users JOIN user_friends ON user_friends.friend_id = users.id AND user_friends.user_id = ? ORDER BY users.id"
	err := repository.DB.Raw(query, userID).Scan(&friends).Error
	return friends, err
}

// DeleteFriend ...
func DeleteFriend(userID, friendID string) error {
	query := "DELETE FROM user_friends WHERE user_id = ? AND friend_id = ? OR user_id = ? AND friend_id = ?;"
	err := repository.DB.Exec(query, userID, friendID, friendID, userID).Error
	return err
}

// CheckFriendByID ...
func CheckFriendByID(userID, friendID string) bool {
	var count int64
	query := "SELECT count(*) FROM user_friends WHERE user_id = ? AND friend_id = ?;"
	repository.DB.Raw(query, userID, friendID).Scan(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}
