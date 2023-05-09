package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func (h *Handlers) Upload(ctx *gin.Context) {
	result := struct {
		Default string `json:"default"`
	}{""}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}

	dir := "/upload/" + time.Now().Format("2006") + "/" + time.Now().Format("0102") + "/"
	err = os.MkdirAll("./public"+dir, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}

	url := dir + uuid.New().String() + strings.ToLower(path.Ext(file.Filename))

	dst := "./public/" + url
	result.Default = url
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
	}

	ctx.JSON(http.StatusOK, result)
}
