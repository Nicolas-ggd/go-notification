// SPDX-License-Identifier: MIT
// Copyright (c) 2024 TOMIOKA
//
// This file is part of the gonotification project.

package repository

import (
	"context"
	"database/sql"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	metakit "github.com/Nicolas-ggd/gorm-metakit"
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
	stm := `INSERT INTO notifications (type, time, message)
	VALUES ($1, $2, $3)
	RETURNING id, type, time, message;`

	err := r.DB.QueryRowContext(context.Background(), stm, model.Type, model.Time, model.Message).
		Scan(&model.ID, &model.Type, &model.Time, &model.Message)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *NotificationRepository) List(meta *metakit.Metadata) (*[]models.Notification, *metakit.Metadata, error) {
	var model []models.Notification

	row := r.DB.QueryRow("SELECT COUNT(*) FROM notifications")
	err := row.Scan(&meta.TotalRows)
	if err != nil {
		return nil, nil, err
	}

	query := "SELECT * FROM notifications"
	rows, err := metakit.SPaginate(r.DB, query, meta)
	if err != nil {
		return nil, nil, err
	}

	meta.SortDirectionParams()
	for rows.Next() {
		var item models.Notification
		model = append(model, item)
	}

	return &model, meta, nil
}
