package controller

import (
	"github.com/go-water/water"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func (*Handlers) Upload(ctx *water.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}

	dir := "/images/" + time.Now().Format("20060102") + "/"
	err = os.MkdirAll("./public"+dir, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}

	url := dir + uuid.New().String() + strings.ToLower(path.Ext(file.Filename))
	dst := "./public/" + url
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
	}

	ctx.JSON(http.StatusOK, water.H{"filePath": url})
}
