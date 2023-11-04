package middleware

import (
	"github.com/arioprima/Point-Of-Sale/config"
	"github.com/arioprima/Point-Of-Sale/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware adalah middleware untuk mengautentikasi dan mengotorisasi pengguna berdasarkan peran.
func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 || fields[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid or missing Authorization header"})
			return
		}

		token = fields[1]

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		secretKey, err := config.LoadConfig(".")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		sub, err := utils.ValidateToken(token, secretKey.TokenSecret)
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
		userID, userIDOk := subObj["user_id"].(string)
		userRole, userRoleOk := subObj["user_role"].(string)

		if !userIDOk || !userRoleOk {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Token sub is missing required properties"})
			return
		}

		// Check user role and authorization
		if !checkAuthorization(userRole, requiredRole) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not authorized"})
			return
		}

		// Set user information in the context
		ctx.Set("currentUser", userID)

		// Continue with the request
		ctx.Next()
	}
}

func checkAuthorization(userRole, requiredRole string) bool {
	switch requiredRole {
	case "admin":
		return userRole == "admin"
	case "staff":
		return userRole == "admin" || userRole == "staff"
	case "employee":
		return userRole == "admin" || userRole == "staff" || userRole == "employee"
	default:
		return false
	}
}
