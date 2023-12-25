package helpers

import (
	"diary_api/database"
	"diary_api/models"
)

func FindEntryById(id string) error {
	var entry models.Entry

	if err := database.Database.Where("id = ?", id).First(&entry).Error; err != nil {
		return err
	}
	return nil

}
