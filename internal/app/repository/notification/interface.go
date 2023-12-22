package notification

import "context"

type NotificationRepo interface {
	NotificationRepoClient
}

type NotificationRepoCommand interface {
	MarkAsReadByID(
		ctx context.Context,
		notifID string,
	) error
}
type NotificationRepoQuery interface {
	GetListNotificationByUserID(
		ctx context.Context,
		userID string,
	) ([]string, error)
}

// To One Signal
type NotificationRepoClient interface {
	SendNotification() error
}
