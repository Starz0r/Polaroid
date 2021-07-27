package routers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gobuffalo/packr"
	echo "github.com/spidernest-go/mux"
	"github.com/spidernest-go/mux/middleware"
)

var TEMPLATES = packr.NewBox("./templates")

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type RespError struct {
	Err string `json:"err"`
}

func ListenAndServe() error {
	tmpl, _ := TEMPLATES.FindString("i.html")
	t := &Template{
		templates: template.Must(template.New("i").Parse(tmpl)),
	}

	r := echo.New()
	r.Renderer = t
	r.BodyLimit(32 * 1024 * 1024) // 32 MB
	r.Use(middleware.Recover(), middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodOptions},
	}))

	r.GET("/:img", getImage)

	v0 := r.Group("/api/v0")

	v0.GET("/user/login", userLogin)
	v0.GET("/user/redirect", userRedirect)

	v0.POST("/image", uploadImage)

	r.Start(":7000")

	return nil
}
