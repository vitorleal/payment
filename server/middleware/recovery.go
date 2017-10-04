package middleware

import (
	"github.com/gin-gonic/gin"
	//"log"
)

// Recovery with default 500 JSON response
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//stack := gin.stack(3)
				//log.Printf("PANIC: %s\n%s", err, stack)
				c.Error(err.(error)).SetMeta("Middleware.Recovery")

				c.AbortWithStatusJSON(500, gin.H{
					"error": err,
				})
			}
		}()

		c.Next()
	}
}
