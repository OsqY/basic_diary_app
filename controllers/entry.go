package controllers

import (
	"diary_api/database"
	"diary_api/helpers"
	"diary_api/models"
	"diary_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEntry(context *gin.Context) {
	var input models.Entry
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

func GetEntryById(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := context.Param("id")
	if err = helpers.FindEntryById(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, entry := range user.Entries {
		if entry.ID == utils.StringToUint(id) {
			context.JSON(http.StatusOK, gin.H{"data": entry})
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{"error": "entry not found"})
}

func DeleteEntry(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := context.Param("id")
	if err = helpers.FindEntryById(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, entry := range user.Entries {
		if entry.ID == utils.StringToUint(id) {
			database.Database.Delete(&entry, id)
			context.JSON(http.StatusOK, gin.H{"data": "entry deleted successfully"})
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{"error": "error deleting your entry"})

}

func UpdateEntry(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")
	if err := helpers.FindEntryById(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, entry := range user.Entries {
		if entry.ID == utils.StringToUint(id) {
			var input models.Entry
			if err := context.ShouldBindJSON(&input); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			database.Database.Model(&models.Entry{}).Where("id = ?", id).Update("content", input.Content)
			context.JSON(http.StatusOK, gin.H{"data": "entry updated successfully"})
			return
		}
	}

	context.JSON(http.StatusBadRequest, gin.H{"error": "error updating your entry"})
}
