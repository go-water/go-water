package router

import (
	. "github.com/go-water/go-water/controller"
	"github.com/kataras/iris/v12"
)

func Start() {
	app := iris.New()
	app.Use(iris.LimitRequestBodySize(1024 * 1024))
	app.HandleDir("/", "./public")
	tmpl := iris.HTML("./views", ".html")
	tmpl.Layout("shared/layout.html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	app.Get("/", H.Index)
	app.Get("/docs", H.ListDoc)
	app.Get("/articles", H.ListArticle)
	app.Get("/article/{id}", H.GetArticle)
	app.Get("/reward", H.Reward)
	app.Get("/about", H.About)

	app.Run(
		iris.Addr(":80"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
