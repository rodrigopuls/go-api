package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse registration id."})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch registration."})
		return
	}

	userId := ctx.GetInt64("userId")

	var register models.Register
	register.EventID = eventId
	register.UserID = userId

	err = register.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create registration."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration created!"})
}

func cancelRegistration(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse registration id."})
		return
	}

	userId := ctx.GetInt64("userId")

	var register models.Register
	register.EventID = eventId
	register.UserID = userId

	err = register.CancelRegistration()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration canceled!"})
}
