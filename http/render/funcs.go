package render

import (
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, contentTpl string, data map[string]interface{}) error {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["path"] = ctx.Path()
	data["app"] = global.App
	return ctx.Render(http.StatusOK, contentTpl, data)
}

func (l *layoutTemplate) Render(w io.Writer, contentTpl string, data interface{}, ctx echo.Context) error {
	layout := "layout.html"
	layoutInter := ctx.Get("layout")
	if layoutInter != nil {
		layout = layoutInter.(string)
	}
	contentTpl = "common/" + layout + "," + contentTpl
	htmlFiles := strings.Split(contentTpl, ",")
	for i, contentTpl := range htmlFiles {
		htmlFiles[i] = global.App.TemplateDir + contentTpl
	}
	tpl, err := template.New(layout).Funcs(funcMap).Funcs(template.FuncMap{"include": tplInclude}).ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	if jsTpl := tpl.Lookup("js"); jsTpl == nil {
		tpl.Parse(`{{define "js"}}{{end}}`)
	}
	if cssTpl := tpl.Lookup("css"); cssTpl == nil {
		tpl.Parse(`{{define "css"}}{{end}}`)
	}
	if seoTpl := tpl.Lookup("seo"); seoTpl == nil {
		tpl.Parse(`{{define "seo"}}
<mate name="keywords" content="` + global.App.SEO["keywords"] + `">
<mate name="description" content="` + global.App.SEO["description"] + `">{{end}}`)
	}
	return tpl.Execute(w, data)
}

func NoNavRender(ctx echo.Context, contentTpl string, data map[string]interface{}) error {
	if data == nil {
		data = make(map[string]interface{})
	}
	ctx.Set("layout", "nova_layout.html")
	return ctx.Render(http.StatusOK, contentTpl, data)
}
