package repository

import (
	"context"
	"database/sql"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
)

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepository(db *sql.DB) NotificationRepository {
	return NotificationRepository{
		DB: db,
	}
}

func (r *NotificationRepository) Insert(model *models.Notification) (*models.Notification, error) {
	stm := `INSERT INTO notifications (type, time, message, is_view)
	VALUES ($1, $2, $3, $4)
	RETURNING id, type, time, message, is_view;`

	err := r.DB.QueryRowContext(context.Background(), stm, model.Type, model.Time, model.Message, model.IsView).
		Scan(&model.ID, &model.Type, &model.Time, &model.Message, &model.IsView)
	if err != nil {
		return nil, err
	}

	return model, nil
}
