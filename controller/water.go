package controller

import (
	"encoding/csv"
	"github.com/kataras/iris/v12"
	"io"
	"os"
)

func (h *Handlers) Reward(ctx iris.Context) {
	csvFile, err := os.Open("./content/reward.csv")
	if err != nil {
		ctx.EndRequest()
	}

	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	result := make([][]string, 0)
	for {
		row, er := csvReader.Read()
		if er == io.EOF {
			break
		} else if err != nil {
			ctx.EndRequest()
		}

		result = append(result, row)
	}

	ctx.ViewData("title", "打赏站长 - 爱斯园")
	ctx.ViewData("body", result)
	ctx.View("reward.html")
}

func (h *Handlers) About(ctx iris.Context) {
	ctx.ViewData("title", "关于网站 - 爱斯园")
	ctx.ViewData("body", 88888)
	ctx.View("about.html")
}
