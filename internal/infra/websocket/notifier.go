package websocket

import "easy-go-monitor/internal/api/handler"

type WebSocketNotifier struct{}

func (n *WebSocketNotifier) PushEvent(eventType string, payload any) {
	handler.PushEvent(eventType, payload)
}
