package ginutil

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

const homeTemplate = `
<html>
<head><title>{{.Title}}</title></head>
<frameset cols="200,*" border="0">
	<frame src="/menu?sn={{.SN}}", name="menu">
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
					<li><a href="{{$v}}?sn={{$.SN}}" target="result">{{$k}}</a></li> 
					{{end}} 
				</ul>
			</body>
		</html>
`

type EmbedRouter struct {
	Menu      string
	RoutePath string
	Handle    gin.HandlerFunc // if nil implement in outer
}

type EmbedDebugWeb struct {
	G        *Engine
	Title    string
	SN       string
	HomePath []string
	Routers  []EmbedRouter
}

func NewEmbedDebugWeb(g *Engine, sn, title string, homePath ...string) *EmbedDebugWeb {
	w := &EmbedDebugWeb{
		G:        g,
		Title:    title,
		SN:       sn,
		HomePath: []string{"/"},
	}

	w.HomePath = append(w.HomePath, homePath...)

	return w
}

func (e *EmbedDebugWeb) AddRouter(menu, routePath string, h gin.HandlerFunc) {
	e.Routers = append(e.Routers, EmbedRouter{
		Menu:      menu,
		RoutePath: routePath,
		Handle:    h,
	})
}

func (e *EmbedDebugWeb) AddFileSystem(menu, routePath, dir string) {
	e.AddRouter(menu, routePath, nil)
	e.G.StaticFS(routePath, gin.Dir(dir, true))
}

func (e *EmbedDebugWeb) Register() {
	tmpl := template.Must(template.New("home").Parse(homeTemplate))
	template.Must(tmpl.New("menu").Parse(menuTemplate))
	e.G.SetHTMLTemplate(tmpl)

	home := func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", map[string]interface{}{
			"Title": e.Title,
			"SN":    e.SN,
		})
	}

	for _, homePath := range e.HomePath {
		e.G.GET(homePath, home)
	}

	data := map[string]interface{}{}

	for _, r := range e.Routers {
		data[r.Menu] = r.RoutePath
		if r.Handle != nil {
			e.G.GET(r.RoutePath, r.Handle)
		}
	}

	e.G.GET("/menu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu", gin.H{
			"Menu": data,
			"SN":   e.SN,
		})
	})
}
