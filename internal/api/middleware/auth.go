package middleware

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/constraints"
	"go-monitor-tool/internal/infra/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(jwtService jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.HandleError(c, constraints.ErrAuthFailed)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := jwtService.ValidateToken(tokenStr)
		if err != nil {
			response.HandleError(c, constraints.ErrAuthFailed)
			return
		}

		c.Set("user_id", userID.String())
		c.Next()
	}
}
