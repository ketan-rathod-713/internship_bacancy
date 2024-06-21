package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Custome Middleware of gin
// it should return gin.HandleFunc
// gin.LoggerWithFormatter is also returns this so we can use it.
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%v - [%v] %v %v %v %v \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
