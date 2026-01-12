package database

import (
	"database/sql"
)

type EventModel struct {
	DB *sql.DB
}

type Events struct {
	ID          int    `json:"id"`
	OwnerID     string `json:"owner_id" binding:"required,min=3"`
	Name        string `json:"name" binding:"required,min=10"`
	Description string `json:"description" binding:"required,min=10"`
	Date        string `json:"date" binding:"required,datetime=2006-01-02"`
	Location    string `json:"location" binding:"required,min=3"`
}
