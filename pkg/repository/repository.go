package repository

import "database/sql"

type Repository struct {
	NotificationRepository NotificationRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NotificationRepository: NewNotificationRepository(db),
	}
}
