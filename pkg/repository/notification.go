package repository

import "database/sql"

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepository(db *sql.DB) NotificationRepository {
	return NotificationRepository{
		DB: db,
	}
}
