package domain

type Notifier interface {
	PushEvent(eventType string, payload any)
}
