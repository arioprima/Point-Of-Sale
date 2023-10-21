package middleware

import (
	"database/sql"
	"github.com/arioprima/blog_web/repository"
	"github.com/arioprima/blog_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func UserHandler(userRepository repository.UserRepository, db *sql.DB, requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 && fields[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid or missing Authorization header"})
			return
		}

		token = fields[1]

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		secretKey := []byte(os.Getenv("JWT_TOKEN_SECRET"))
		sub, err := utils.ValidateToken(token, string(secretKey))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		// Check if "sub" is a valid JSON object
		subObj, ok := sub.(map[string]interface{})
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Token sub is not a valid object"})
			return
		}

		// Access the properties within "sub"
		userID, userIDOk := subObj["id"].(string)

		// berdasarkan roke
		userRole, userRoleOk := subObj["role"].(string)

		if !userIDOk && !userRoleOk {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Token sub is missing required properties"})
			return
		}

		//role harus admin
		if userRole != requiredRole {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not authorized"})
			return
		}

		if userRepository == nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "User repository is nil"})
			return
		}

		// Begin a new transaction
		tx, err := db.Begin()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error starting a transaction"})
			return
		}

		// Handle the error returned by FindByUserName
		result, err := userRepository.FindById(ctx, tx, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.Set("currentUser", result.ID)
		ctx.Next()
	}
}
