package utils

import (
	"github.com/gin-gonic/gin"
)

// Response is the server response
type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Abort replies with an error
func Abort(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, Response{
		Error:   true,
		Message: message,
	})
}

// Reply replies with success
func Reply(c *gin.Context, code int, message string, result interface{}) {
	c.JSON(code, Response{
		Error:   false,
		Message: message,
		Result:  result,
	})
}
