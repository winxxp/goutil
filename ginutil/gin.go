package ginutil

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Name() string {
	return "ginutil"
}

type Engine struct {
	*gin.Engine
	logger ILogger
}

func NewGin(logger ILogger) *Engine {
	engine := &Engine{
		Engine: gin.New(),
		logger: logger,
	}

	engine.Use(Logger(logger), Recovery(logger))

	return engine
}

func (e *Engine) Run(ctx context.Context, addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: e.Engine,
	}

	go func() {
		<-ctx.Done()
		err := server.Close()
		e.logger.Error(fmt.Sprintf("server close: %v", err))
	}()

	return server.ListenAndServe()
}

type ILogger interface {
	Info(i string)
	Error(e string)
}
