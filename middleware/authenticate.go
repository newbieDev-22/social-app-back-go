package middleware

import (
	"net/http"
	"simple-social-app/db"
	"simple-social-app/dto"
	"simple-social-app/model"
	"simple-social-app/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		accessToken, err := jwtService.ValidateToken(authHeader)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		if !accessToken.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		userId, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		var user model.User
		err = db.Conn.Where("id = ?", userId).First(&user).Error
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
			return
		}

		ctx.Set("user", dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
		ctx.Next()
	}
}
