package config

const (
	StreamName = "NOTIFICATION"

	SubjectBroadcastNotification = "NOTIFICATION.send-to-all"
	SubjectClientNotification    = "NOTIFICATION.send-to-clients"
	SubjectNotificationList      = "NOTIFICATION.list"
	SubjectVersion               = "0.0.1"
)

const (
	MigrationURL = "file://internal/migrations"
	DatabaseName = "sqlite3"
	DBName       = "notifications"
)
