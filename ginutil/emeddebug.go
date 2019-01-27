package ginutil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"net/url"
)

const homeTemplate = `
<html>
<head><title>{{.Title}}</title></head>
<frameset cols="200,*" border="0">
	<frame src="/menu?{{.Param}}", name="menu">
	<frame src="/menu", name="result" >
</frameset>
</html>
`
const menuTemplate = `
		<html>
			<head><head>
			<body>
				<h3>Menu</h3>
				<ul>
					{{range $k,$v := .Menu}} 
					<li><a href="{{$v}}" target="result">{{$k}}</a></li> 
					{{end}} 
				</ul>
			</body>
		</html>
`

type EmbedRouter struct {
	Menu      string
	RoutePath string
	Param     gin.H           // url parameters
	Handle    gin.HandlerFunc // if nil implement in outer
}

type EmbedDebugWeb struct {
	G        *Engine
	Title    string
	Param    gin.H
	HomePath []string
	Routers  []EmbedRouter
}

func NewEmbedDebugWeb(g *Engine, title string, defParam gin.H, homePath ...string) *EmbedDebugWeb {
	w := &EmbedDebugWeb{
		G:        g,
		Title:    title,
		Param:    defParam,
		HomePath: []string{"/"},
	}

	w.HomePath = append(w.HomePath, homePath...)

	return w
}

func (e *EmbedDebugWeb) AddRouter(menu, routePath string, param gin.H, h gin.HandlerFunc) {
	e.Routers = append(e.Routers, EmbedRouter{
		Menu:      menu,
		RoutePath: routePath,
		Param:     param,
		Handle:    h,
	})
}

func (e *EmbedDebugWeb) AddFileSystem(menu, routePath, dir string, param gin.H) {
	e.AddRouter(menu, routePath, param, nil)
	e.G.StaticFS(routePath, gin.Dir(dir, true))
}

func (e *EmbedDebugWeb) Register() {
	tmpl := template.Must(template.New("home").Parse(homeTemplate))
	template.Must(tmpl.New("menu").Parse(menuTemplate))
	e.G.SetHTMLTemplate(tmpl)

	homeUrlParam := url.Values{}
	for k, v := range e.Param {
		homeUrlParam.Set(k, fmt.Sprint(v))
	}

	home := func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", map[string]interface{}{
			"Title": e.Title,
			"Param": homeUrlParam.Encode(),
		})
	}

	for _, homePath := range e.HomePath {
		e.G.GET(homePath, home)
	}

	data := map[string]interface{}{}

	for _, r := range e.Routers {
		routeUrlParam := url.Values{}

		for k, v := range r.Param {
			routeUrlParam.Set(k, fmt.Sprint(v))
		}
		for k, v := range e.Param {
			routeUrlParam.Set(k, fmt.Sprint(v))
		}

		data[r.Menu] = r.RoutePath + "?" + routeUrlParam.Encode()
		if r.Handle != nil {
			e.G.GET(r.RoutePath, r.Handle)
		}
	}

	e.G.GET("/menu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu", gin.H{
			"Menu": data,
		})
	})
}
