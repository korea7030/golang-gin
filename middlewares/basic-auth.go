package middlewares

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	// gin.BasicAuth Return
	return gin.BasicAuth(gin.Accounts{
		"pragmatic": "reviews",
	})
}
