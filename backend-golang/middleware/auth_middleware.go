package middleware

import (
	"database/sql"
	"github.com/arioprima/Point-Of-Sale/config"
	"github.com/arioprima/Point-Of-Sale/repository"
	"github.com/arioprima/Point-Of-Sale/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UserHandler(userRepository repository.UserRepository, db *sql.DB, requiredRole string) gin.HandlerFunc {
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

		// berdasarkan roke
		userRole, userRoleOk := subObj["user_role"].(string)

		//check error for userIDOk && userRoleOk

		/* ini kode untuk cari error
		log.Println("sub", sub)
		log.Println("user_id", userIDOk)
		log.Println("user_role", userRoleOk)
		*/

		if !userIDOk && !userRoleOk {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Token sub is missing required properties"})
			return
		}

		//role harus admin
		if (requiredRole == "admin" && userRole != "admin") ||
			(requiredRole == "staff" && userRole != "admin" && userRole != "staff") ||
			(requiredRole == "employee" && userRole != "admin" && userRole != "staff" && userRole != "employee") {
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
