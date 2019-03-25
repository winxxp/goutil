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
	if err != nil {
		c.AbortWithStatusJSON(code, map[string]interface{}{
			"code":    code,
			"message": err.Error(),
		})
		c.Error(err)
	} else {
		c.AbortWithStatus(code)
	}
}

func (c *Context) Uint64(key string, getter Getter) (u uint64) {
	if c.err != nil {
		return
	}

	u, c.err = strconv.ParseUint(getter.Get(key), 10, 0)
	return
}

func (c *Context) Uint64FromParam(key string) uint64 {
	return c.Uint64(key, c.ParamGetter())
}

func (c *Context) Uint64FromQuery(key string, defaultValue ...uint64) uint64 {
	getter := c.QueryGetter()
	if len(defaultValue) > 0 {
		getter = c.DefaultQueryGetter(strconv.FormatUint(defaultValue[0], 10))
	}

	return c.Uint64(key, getter)
}

func (c *Context) Int64(key string, getter Getter) (i int64) {
	if c.err != nil {
		return
	}

	i, c.err = strconv.ParseInt(getter.Get(key), 10, 0)
	return
}

func (c *Context) Int64FromParam(key string) int64 {
	return c.Int64(key, c.ParamGetter())
}

func (c *Context) Int64FromQuery(key string, defaultValue ...int64) int64 {
	getter := c.ParamGetter()
	if len(defaultValue) > 0 {
		getter = c.DefaultQueryGetter(strconv.FormatInt(defaultValue[0], 10))
	}

	return c.Int64(key, getter)
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

func (c *Context) StrFromParam(key string) string {
	return c.Str(key, c.ParamGetter())
}

//StrFromQuery 从Path query中取值，值不能为空
func (c *Context) StrFromQuery(key string, defaultValue ...string) string {
	getter := c.QueryGetter()
	if len(defaultValue) > 0 {
		getter = c.DefaultQueryGetter(defaultValue[0])
	}

	return c.Str(key, getter)
}

func (c *Context) Time(key string, layout string, getter Getter) (t time.Time) {
	if c.err != nil {
		return
	}

	t, c.err = time.Parse(layout, getter.Get(key))
	return
}

func (c *Context) TimeFromParam(key string, layout string) (t time.Time) {
	return c.Time(key, layout, c.ParamGetter())
}

func (c *Context) TimeFromQuery(key string, layout string) (t time.Time) {
	return c.Time(key, layout, c.QueryGetter())
}

func (c *Context) QueryGetter() GetHandle {
	return func(key string) string {
		return c.Context.Query(key)
	}
}

func (c *Context) DefaultQueryGetter(def string) GetHandle {
	return func(key string) string {
		return c.Context.DefaultQuery(key, def)
	}
}

func (c *Context) ParamGetter() GetHandle {
	return func(key string) string {
		return c.Context.Param(key)
	}
}

func (c *Context) RangeWithDefault() (offset, points int64) {
	return c.Range(0, -1)
}

func (c *Context) Range(o, p int) (offset, points int64) {
	offset = c.Int64("offset", c.DefaultQueryGetter(strconv.Itoa(o)))
	points = c.Int64("points", c.DefaultQueryGetter(strconv.Itoa(p)))
	return
}

func (c *Context) PageWithDefault() (offset, count uint64) {
	return c.Page(0, 100)
}

func (c *Context) Page(o, cnt int) (offset, count uint64) {
	offset = c.Uint64("offset", c.DefaultQueryGetter(strconv.Itoa(o)))
	count = c.Uint64("count", c.DefaultQueryGetter(strconv.Itoa(cnt)))
	return
}
