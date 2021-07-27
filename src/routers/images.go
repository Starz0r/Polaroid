package routers

import (
	"io/ioutil"
	mime_stdlib "mime"
	"net/http"
	"os"

	"github.com/Starz0r/Polaroid/src/crypto"
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

	// heuristics
	imgfile, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, RespError{Err: "image was corrupt or missing"})
	}

	file, err := imgfile.Open()
	defer file.Close()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, RespError{Err: "file could not be opened"})
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, RespError{Err: "file could not be read"})
	}
	mime := http.DetectContentType(buf)
	mimefext, err := mime_stdlib.ExtensionsByType(mime)
	if err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, RespError{Err: "media type did not have a valid extensions"})
	}

	// start the transfer
	filename := *new(string)
	for {
		filename = crypto.String(6) + mimefext[0]
		if !objstore.IsNameConflicting(filename) {
			break
		}
	}
	err = objstore.Upload(file, filename, "public-read", mime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RespError{Err: "transfer failed"})
	}

	return c.JSON(http.StatusAccepted, Image{URL: PUBLICURL + imgfile.Filename})
}

func getImage(c echo.Context) error {
	img := c.Param("img")

	resp, err := http.Get(S3PUBLICURL + img)
	if err != nil {
		return c.JSON(http.StatusNotFound, RespError{Err: "no image"})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RespError{Err: "file could not be opened"})
	}
	mime := http.DetectContentType(body)

	return c.Blob(http.StatusOK, mime, body)
}
