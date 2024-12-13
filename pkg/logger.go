package pkg

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	DefaultLogFormat = "%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n"
	DefaultLogFile   = "app.log"
)

func CustomLogger() (io.Writer, gin.HandlerFunc) {
	logFile, _ := os.Create(DefaultLogFile)
	customWriter := io.MultiWriter(logFile, gin.DefaultWriter)

	return customWriter, gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(DefaultLogFormat,
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
