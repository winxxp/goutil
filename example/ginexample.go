package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/winxxp/goutil"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	ctx, _ := context.WithCancel(context.Background())

	g := goutil.NewGin(&Log{})

	emb := goutil.NewEmbedDebugWeb(g, "Gin Example", gin.H{"sn": 1234}, "/home")
	emb.AddRouter("Config", "/config", gin.H{"p1": 1, "p2": 2}, func(c *gin.Context) {
		param := gin.H{}
		for k, v := range c.Request.URL.Query() {
			param[k] = v
		}

		c.JSON(http.StatusOK, gin.H{
			"param": param,
		})
	})
	emb.Register()

	log.Println("addr:", ":8080", "server run")
	err := g.Run(ctx, ":18080")
	log.Println(err, "server quit")
}

type Log struct {
}

func (Log) Info(i string) {
	log.Println(i)
}

func (Log) Error(e string) {
	log.Println(e)
}
