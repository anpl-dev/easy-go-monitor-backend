package handler

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// --- WebSocket設定 ---
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.RWMutex
	broadcast = make(chan map[string]any, 100) // イベント送信用チャネル
)

// --- WebSocket接続ハンドラ ---
func WebsocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "WebSocket upgrade failed")
		return
	}

	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
		conn.Close()
	}()

	// クライアントからのメッセージは無視
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

// --- イベント配信用ループ ---
func BroadcastEvents() {
	for msg := range broadcast {
		clientsMu.RLock()
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				client.Close()
				clientsMu.RUnlock()
				clientsMu.Lock()
				delete(clients, client)
				clientsMu.Unlock()
				clientsMu.RLock()
			}
		}
		clientsMu.RUnlock()
	}
}

// --- 外部からイベントを送信 ---
func PushEvent(eventType string, data any) {
	broadcast <- map[string]any{
		"type": eventType,
		"data": data,
	}
}
