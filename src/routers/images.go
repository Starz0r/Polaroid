package routers

import (
	"net/http"
	"os"

	"github.com/Starz0r/Polaroid/src/database"
	"github.com/Starz0r/Polaroid/src/objstore"
	echo "github.com/spidernest-go/mux"
)

var S3PUBLICURL = os.Getenv("S3_PUBLIC_URL")
var PUBLICURL = os.Getenv("PUBLIC_URL")

type Image struct {
	URL string `json:"url"`
}

func uploadImage(c echo.Context) error {
	// check if authorized
	req := c.Request()
	key := string(req.Request.Header.Peek("App-Key"))
	if key == "" {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "no appkey in header"})
	}
	appkey := &database.AppKey{Key: key}
	exists, err := appkey.KeyExists()
	if exists == false || err != nil {
		return c.JSON(http.StatusUnauthorized, RespError{Err: "invalid appkey"})
	}

	// start the transfer
	imgfile, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, RespError{Err: "image was corrupt or missing"})
	}

	file, err := imgfile.Open()
	defer file.Close()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, RespError{Err: "file could not be opened"})
	}

	err = objstore.Upload(file, imgfile.Filename, "public-read")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RespError{Err: "transfer failed"})
	}

	return c.JSON(http.StatusAccepted, Image{URL: PUBLICURL + imgfile.Filename})
}

func getImage(c echo.Context) error {
	img := c.Param("img")
	return c.Render(http.StatusOK, "i", Image{URL: S3PUBLICURL + img})
}
