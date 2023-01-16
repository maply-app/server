package managers

import (
	"maply/models"
	"maply/repository"
)

// GetRequestByID ...
func GetRequestByID(requestID string) (models.Request, error) {
	var request models.Request
	result := repository.DB.Preload("Sender").First(&request, "id = ?", requestID)
	return request, result.Error
}

// DeleteRequestByID ...
func DeleteRequestByID(requestID string) (models.Request, error) {
	var request models.Request
	result := repository.DB.Where("id = ?", requestID).Delete(&request)
	return request, result.Error
}

// DeleteRequest ...
func DeleteRequest(userID, requestID string) error {
	query := "DELETE FROM requests WHERE id = ? AND sender_id = ? OR receiver_id = ?;"
	err := repository.DB.Exec(query, requestID, userID, userID).Error
	return err
}

// FindRequestBySenderAndReceiver ...
func FindRequestBySenderAndReceiver(senderID, receiverID string) (models.Request, error) {
	var request models.Request
	query := "SELECT * FROM requests WHERE sender_id = ? AND receiver_id = ? OR sender_id = ? AND receiver_id = ?;"
	err := repository.DB.Raw(query, senderID, receiverID, receiverID, senderID).Scan(&request).Error
	return request, err
}

// GetRequestsByReceiver ...
func GetRequestsByReceiver(userID string) ([]models.Request, error) {
	var requests []models.Request
	query := "SELECT * FROM requests WHERE receiver_id = ?;"
	err := repository.DB.Raw(query, userID).Preload("Sender").Find(&requests).Error
	return requests, err
}

// GetRequestsBySender ...
func GetRequestsBySender(userID string) ([]models.Request, error) {
	var requests []models.Request
	query := "SELECT * FROM requests WHERE sender_id = ?;"
	err := repository.DB.Raw(query, userID).Preload("Receiver").Find(&requests).Error
	return requests, err
}

// CreateRequest ...
func CreateRequest(r *models.Request) (string, error) {
	result := repository.DB.Create(&r)
	return r.ID, result.Error
}
