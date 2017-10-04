package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ingresse/payment/errors"
	g "github.com/ingresse/payment/gateway"
	"strings"
)

// Jwt validates the Authorization header and the jwt token
// with the companyId and userId fields
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"]

		// Validate that Authorization header is present
		if authHeader == nil {
			e := errors.UnauthorizedError(1002, "Missing Authorization header")
			c.AbortWithStatusJSON(e.HttpStatus, e)
			return
		}

		jwt := strings.Split(strings.Replace(authHeader[0], "Bearer ", "", 1), ".")

		// Validatei that the token has user data information
		if len(jwt) < 2 {
			e := errors.UnauthorizedError(1003, "Invalid Authorization token")
			c.AbortWithStatusJSON(e.HttpStatus, e)
			return
		}

		// Decode jwt user data
		loggedUser := &g.LoggedUser{}
		decoded, _ := base64.StdEncoding.DecodeString(jwt[1] + "==")

		// Convert the data into a LoggedUser data
		err := json.Unmarshal(decoded, &loggedUser)

		// Validate user data
		if err != nil || loggedUser.IsNotValid() {
			e := errors.UnauthorizedError(1004, "Invalid Authorization token data")
			c.AbortWithStatusJSON(e.HttpStatus, e)
			return
		}

		c.Set("loggedUser", loggedUser)
		c.Next()
	}
}
