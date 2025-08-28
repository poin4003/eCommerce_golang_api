package middlewares

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poin4003/eCommerce_golang_api/internal/consts"
	"github.com/poin4003/eCommerce_golang_api/internal/utils/auth"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the request url path
		uri := c.Request.URL.Path
		log.Println("uri request: ", uri)
		// Check headers authrization
		jwtToken, valid := auth.ExtractBearerToken(c)
		if !valid {
			c.AbortWithStatusJSON(401, gin.H{
				"code":        40001,
				"err":         "Unauthorized",
				"description": "",
			})
			return
		}

		// validate jwt token by subject
		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"code":        40001,
				"err":         "Invalid token",
				"description": "",
			})
			return
		}

		// update claims to context
		ctx := context.WithValue(c.Request.Context(), consts.SUBJECT_UUID_KEY, claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
