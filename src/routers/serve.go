package routers

import (
	echo "github.com/spidernest-go/mux"
	"github.com/spidernest-go/mux/middleware"
)

type RespError struct {
	Err string `json:"err"`
}

func ListenAndServe() error {
	r := echo.New()
	r.BodyLimit(32 * 1024 * 1024) // 32 MB
	r.Use(middleware.Recover())

	v0 := r.Group("/api/v0")

	v0.GET("/user/login", userLogin)
	v0.GET("/user/redirect", userRedirect)

	v0.POST("/image", uploadImage)

	r.Start(":7000")

	return nil
}
