package ginutil

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleJSONString(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		out, _ := json.MarshalIndent(data, "", "    ")
		c.String(http.StatusOK, string(out))
	}
}
