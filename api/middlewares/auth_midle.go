package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/web-gopro/auth_exam/token"
)

// func AuthMiddlewareSuperAdmin() gin.HandlerFunc {

// 	return func(ctx *gin.Context) {
// 		tokenString := ctx.GetHeader("authorization")

// 		if tokenString == "" {
// 			ctx.JSON(401, gin.H{"error": "authorization token not provided"})
// 			ctx.Abort()
// 		}

// 		claim, err := token.ParseJWT(tokenString)
// 		if err != nil {
// 			ctx.JSON(401, gin.H{"error": err.Error()})
// 			ctx.Abort()
// 		}

// 		if claim.UserRole != "superadmin" {
// 			ctx.JSON(401, gin.H{"error": "your role isn't superadmin "})
// 			ctx.Abort()
// 		}

//			ctx.Next()
//		}
//	}
func AuthMiddlewareSuperAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("authorization")

		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "authorization token not provided"})
			ctx.Abort()
			return
		}

		claim, err := token.ParseJWT(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return // <--- BU YERDA RETURN QO‘SHISH KERAK
		}

		if claim == nil || claim.UserRole != "superadmin" {
			ctx.JSON(401, gin.H{"error": "your role isn't superadmin"})
			ctx.Abort()
			return // <--- BU YERDA HAM RETURN KERAK
		}

		ctx.Next()
	}
}
