package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		fmt.Printf("%+v\n", params)
		return fmt.Sprintf(`%s - [%s] %s %s %d %s`,
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
