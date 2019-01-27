package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/winxxp/glog"
	"github.com/winxxp/goutil/ginutil"
	"net/http"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	g := ginutil.NewGin()

	emb := ginutil.NewEmbedDebugWeb(g, "Gin Example", gin.H{"sn": 1234}, "/home")
	emb.AddRouter("Config", "/config", gin.H{"p1": 1, "p2":2}, func(c *gin.Context) {
		param := gin.H{}
		for k, v := range c.Request.URL.Query() {
			param[k] = v
		}

		c.JSON(http.StatusOK, gin.H{
			"param": param,
		})
	})
	emb.Register()

	glog.WithField("addr:", ":8080").Info("server run")
	err := g.Run(ctx, ":8080")
	glog.WithResult(err).Log("server quit")
}
