package main

import (
	"app-event/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Events

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := app.models.Events.Insert(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create event.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    event,
		"message": "Event created successfully.",
	})
}

func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	event, err := app.models.Events.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve event.",
		})
		return
	}

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": event,
	})
}

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	existingEvent, err := app.models.Events.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve event.",
		})
		return
	}

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found.",
		})
		return
	}

	var updatedEvent database.Events
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedEvent.ID = id

	if err := app.models.Events.Update(&updatedEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update event.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event successfully updated",
		"data":    updatedEvent,
	})
}

func (app *application) getAllEvents(c *gin.Context) {
	events, err := app.models.Events.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve events.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	event, err := app.models.Events.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve event.",
		})
		return
	}

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found.",
		})
		return
	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete event.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully.",
	})
}
