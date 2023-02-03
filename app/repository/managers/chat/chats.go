package chat

import (
	"maply/models"
	"maply/repository"
)

func GetChats(userId string, count, offset int) ([]models.Message, error) {
	var messages []models.Message
	query := `SELECT m1.sender_id, m1.receiver_id, m1.text, m1.created_at
			  FROM messages m1
			  JOIN ( SELECT LEAST(sender_id, receiver_id) user1,
							GREATEST(sender_id, receiver_id) user2,
							MAX(created_at) created_at
					   FROM messages m2 WHERE m2.sender_id = ? OR m2.receiver_id = ?
					   GROUP BY user1, user2 ) m3 ON m1.sender_id IN (m3.user1, m3.user2)
												  AND m1.sender_id IN (m3.user1, m3.user2)
												  AND m1.created_at = m3.created_at
			  									  ORDER BY m1.created_at DESC OFFSET ? LIMIT ?;`
	err := repository.DB.Raw(query, userId, userId, offset, count).Preload(
		"Sender",
	).Preload(
		"Receiver",
	).Find(&messages).Error
	return messages, err
}
