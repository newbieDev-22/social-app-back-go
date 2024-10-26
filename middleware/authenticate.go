package middleware

import (
	"net/http"
	"simple-social-app/dto"
	"simple-social-app/repository"
	"simple-social-app/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService, userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		accessToken, err := jwtService.ValidateToken(authHeader)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
			return
		}

		if !accessToken.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
			return
		}

		userId, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
			return
		}

		user, err := userRepository.GetUserById(ctx, nil, userId)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
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
