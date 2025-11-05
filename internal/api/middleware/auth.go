package middleware

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/infra/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(jwtService jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.HandleError(c, codes.ErrAuthFailed)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := jwtService.ValidateToken(tokenStr)
		if err != nil {
			response.HandleError(c, codes.ErrAuthFailed)
			return
		}

		c.Set("user_id", userID.String())
		c.Next()
	}
}
