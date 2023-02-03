package chat

import (
	"maply/models"
	"maply/repository"
)

func CreateMessage(r *models.Message) error {
	return repository.DB.Create(&r).Error
}

func GetMessages(userId, receiverID string, count, offset int) ([]models.Message, error) {
	var messages []models.Message
	query := `SELECT * FROM messages WHERE
                           sender_id = ? AND receiver_id = ? OR
                           sender_id = ? AND receiver_id = ?
                       ORDER BY created_at DESC OFFSET ? LIMIT ?;`
	err := repository.DB.Raw(
		query,
		userId,
		receiverID,
		receiverID,
		userId,
		offset,
		count,
	).Preload("Sender").Find(&messages).Error
	return messages, err
}

func ReadMessages(senderId, receiverID string) error {
	query := `UPDATE 
			    messages 
			  SET 
			    is_read = true 
			  WHERE
			    sender_id = ?
			    AND receiver_id = ?
			    AND NOT is_read;`
	err := repository.DB.Exec(
		query,
		senderId,
		receiverID,
	).Error
	return err
}

func GetUnreadMessages(senderId, receiverID string) (int, error) {
	var count int
	query := `SELECT 
			    count(*) 
			  FROM 
			    messages 
			  WHERE 
			    sender_id = ?
			    AND receiver_id = ?
			    AND NOT is_read;`
	err := repository.DB.Raw(
		query,
		senderId,
		receiverID,
	).Find(&count).Error
	return count, err
}
