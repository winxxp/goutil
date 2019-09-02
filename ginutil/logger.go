package ginutil

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	IDGen *snowflake.Node
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

func init() {
	var err error
	IDGen, err = snowflake.NewNode(time.Now().Unix() % 1024)
	if err != nil {
		panic(err)
	}
}

func Logger(logger ILogger, noLog ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(noLog); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range noLog {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {

		start := time.Now()
		rid := IDGen.Generate().Base58()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Set("rid", rid)

		if logger == nil {
			return
		}

		c.Next()

		if _, ok := skip[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			entry, _ := json.Marshal(gin.H{
				"rid":     rid,
				"latency": latency,
				"client":  clientIP,
				"comment": comment,
				"raw":     raw,
			})

			var logFn func(s string)

			switch statusCode / 100 {
			case 1, 2, 3:
				logFn = logger.Info
			default:
				logFn = logger.Error
			}

			logFn(fmt.Sprintf("%-7s%-4d%s %+v", method, statusCode, path, string(entry)))
		}
	}
}
