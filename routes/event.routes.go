package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		log.Println("error getting events", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error getting events", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println("error parsing id", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error parsing id", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		log.Println("error getting event", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "error getting event", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func postEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error"})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := ctx.MustGet("claims").(*utils.ParsedClaims)
	event.UserID = claims.UserID
	err = event.Save()
	if err != nil {
		log.Println("error saving event", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error saving event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println("error parsing id", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error parsing id", "error": err.Error()})
		return
	}
	event , err := models.GetEventByID(eventId)
	if err != nil {
		log.Println("error getting event", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "event not found", "error": err.Error()})
		return
	}
	if event.UserID != ctx.MustGet("claims").(*utils.ParsedClaims).UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}
	var eventUpdate models.Event

	err = ctx.ShouldBindJSON(&eventUpdate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "error parsing event data"})
		return
	}

	eventUpdate.ID = eventId
	eventUpdate.UserID = event.UserID
	err = eventUpdate.Update()
	if err != nil {
		log.Println("error updating event", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error updating event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event updated successfully", "event": eventUpdate})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println("error parsing id", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error parsing id", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		log.Println("error getting event", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "event not found", "error": err.Error()})
		return
	}
	if event.UserID != ctx.MustGet("claims").(*utils.ParsedClaims).UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		log.Println("error deleting event", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting event", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
