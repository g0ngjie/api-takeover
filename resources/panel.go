package resources

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed panel/*
var panelFS embed.FS

//go:embed assets/*
var assetsFS embed.FS

// 静态资源配置
func Init(app *gin.Engine) {
	// 引入js、css等
	app.Any("/assets/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(assetsFS))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
	// 引入html
	app.SetHTMLTemplate(template.Must(template.New("").ParseFS(panelFS, "panel/*")))

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
