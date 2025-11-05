package middleware

import (
	"easy-go-monitor/internal/infra/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleWare(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // execute handler

		duration := time.Since(start)

		status := c.Writer.Status()

		log.Info("HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", status,
			"latency_ms", duration.Milliseconds(),
			"clitent_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
		)
	}
}
