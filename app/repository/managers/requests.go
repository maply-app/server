package managers

import (
	"maply/models"
	"maply/repository"
)

func GetRequestByID(requestID string) (models.Request, error) {
	var request models.Request
	result := repository.DB.Preload("Sender").First(&request, "id = ?", requestID)
	return request, result.Error
}

func DeleteRequestByID(requestID string) (models.Request, error) {
	var request models.Request
	result := repository.DB.Where("id = ?", requestID).Delete(&request)
	return request, result.Error
}

func DeleteRequest(userId, requestID string) error {
	query := "DELETE FROM requests WHERE id = ? AND sender_id = ? OR receiver_id = ?;"
	err := repository.DB.Exec(query, requestID, userId, userId).Error
	return err
}

func FindRequestBySenderAndReceiver(senderID, receiverID string) (models.Request, error) {
	var request models.Request
	query := "SELECT * FROM requests WHERE sender_id = ? AND receiver_id = ? OR sender_id = ? AND receiver_id = ?;"
	err := repository.DB.Raw(query, senderID, receiverID, receiverID, senderID).Scan(&request).Error
	return request, err
}

func GetRequestsByReceiver(userId string) ([]models.Request, error) {
	var requests []models.Request
	query := "SELECT * FROM requests WHERE receiver_id = ?;"
	err := repository.DB.Raw(query, userId).Preload("Sender").Find(&requests).Error
	return requests, err
}

func GetRequestsBySender(userId string) ([]models.Request, error) {
	var requests []models.Request
	query := "SELECT * FROM requests WHERE sender_id = ?;"
	err := repository.DB.Raw(query, userId).Preload("Receiver").Find(&requests).Error
	return requests, err
}

func CreateRequest(r *models.Request) (string, error) {
	result := repository.DB.Create(&r)
	return r.ID, result.Error
}
