package ginutil

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/winxxp/glog"
	"net/http"
	"strconv"
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

func Logger(notlogged ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		start := time.Now()
		rid := IDGen.Generate().Base58()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Set("rid", rid)
		c.Next()

		if _, ok := skip[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			var entry = glog.WithID(rid).Depth(1).WithFields(glog.Fields{
				"latency": latency,
				"client":  clientIP,
				"comment": comment,
				"raw":     raw,
			})

			var logPad func(s string, rs string, pad byte)
			switch {
			case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices:
				logPad = entry.PadInfo
			case statusCode >= http.StatusMultipleChoices && statusCode < http.StatusBadRequest:
				logPad = entry.PadWarning
			case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
				logPad = entry.PadError
			default:
				logPad = entry.PadError
			}

			logPad(fmt.Sprintf("%-7s%s", method, path), strconv.Itoa(statusCode), '-')
		}
	}
}
