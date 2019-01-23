package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"strconv"
	"time"
)

type Context struct {
	*gin.Context
	err error
}

type Getter interface {
	Get(string) string
}

type PostFormHandle func(key string) (string, bool)

func (p PostFormHandle) Get(key string) string {
	v, _ := p(key)
	return v
}

type GetHandle func(key string) string

func (p GetHandle) Get(key string) string {
	return p(key)
}

func NewContext(c *gin.Context) *Context {
	return &Context{
		Context: c,
	}
}

func (c *Context) Err() error {
	return c.err
}

func (c *Context) AbortWithErrorJSON(code int, err error) {
	c.AbortWithStatusJSON(code, map[string]interface{}{
		"code":    code,
		"message": err.Error(),
	})
	c.Error(err)
}

func (c *Context) Uint64(key string, getter Getter) (u uint64) {
	if c.err != nil {
		return
	}

	u, c.err = strconv.ParseUint(getter.Get(key), 10, 0)
	return
}

func (c *Context) Int64(key string, getter Getter) (i int64) {
	if c.err != nil {
		return
	}

	i, c.err = strconv.ParseInt(getter.Get(key), 10, 0)
	return
}

func (c *Context) Str(key string, getter Getter) (str string) {
	if c.err != nil {
		return
	}

	if str = getter.Get(key); len(str) == 0 {
		c.err = errors.Errorf("%s is empty", key)
	}

	return
}

func (c *Context) Time(key string, layout string, getter Getter) (t time.Time) {
	if c.err != nil {
		return
	}

	t, c.err = time.Parse(layout, getter.Get(key))
	return
}
