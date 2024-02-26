package router

import (
	. "github.com/go-water/go-water/controller"
	"github.com/go-water/go-water/middleware"
	"github.com/go-water/water"
	"github.com/go-water/water/multitemplate"
)

func Start() {
	router := water.New()
	router.Static("/styles", "./public/styles")
	router.Static("/scripts", "./public/scripts")
	router.Static("/images", "./public/images")
	router.StaticFile("/banner.png", "./public/banner.png")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.StaticFile("/pay.jpg", "./public/pay.jpg")
	router.StaticFile("/list.png", "./public/list.png")
	router.StaticFile("/placeholder.jpg", "./public/placeholder.jpg")
	router.HTMLRender = createMyRender()

	router.GET("/", H.Index)
	router.GET("/docs", H.ListDoc)
	router.GET("/doc/{id}", H.GetDoc)
	router.GET("/reward", H.Reward)
	router.GET("/about", H.About)

	router.GET("/login", H.Login)
	router.POST("/login", H.LoginPost)
	router.GET("/logout", H.Logout)

	admin := router.Group("/admin")
	{
		admin.Use(middleware.CheckAuth)
		admin.POST("/upload", H.Upload)
		admin.GET("/add", H.Add)
		admin.POST("/add", H.AddPost)
		admin.GET("/update/{id}", H.Update)
		admin.POST("/update", H.UpdatePost)
		admin.GET("/list", H.List)
	}

	router.Serve(":80")
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "views/shared/layout.html", "views/index.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("docs", "views/shared/layout.html", "views/docs.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("detail", "views/shared/layout.html", "views/detail.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("about", "views/shared/layout.html", "views/about.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("reward", "views/shared/layout.html", "views/reward.html", "views/shared/_header.html", "views/shared/_footer.html")

	r.AddFromFiles("login", "views/auth/login.html")
	r.AddFromFiles("list", "views/admin/layout.html", "views/admin/list.html", "views/admin/_header.html", "views/admin/_footer_none.html")
	r.AddFromFiles("add", "views/admin/layout.html", "views/admin/add.html", "views/admin/_header.html", "views/admin/_footer.html")
	r.AddFromFiles("update", "views/admin/layout.html", "views/admin/update.html", "views/admin/_header.html", "views/admin/_footer.html")
	return r
}
