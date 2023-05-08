package router

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	. "github.com/go-water/go-water/controller"
	"github.com/go-water/go-water/middleware"
)

func Start() {
	router := gin.Default()
	router.Static("/styles", "./public/styles")
	router.Static("/scripts", "./public/scripts")
	router.Static("/images", "./public/images")
	router.StaticFile("/banner.png", "./public/banner.png")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.StaticFile("/pay.jpg", "./public/pay.jpg")
	router.StaticFile("/placeholder.jpg", "./public/placeholder.jpg")
	router.HTMLRender = createMyRender()

	router.GET("/", H.Index)
	router.GET("/docs", H.ListDoc)
	router.GET("/articles", H.ListArticle)
	router.GET("/article/:id", H.GetArticle)
	router.GET("/reward", H.Reward)
	router.GET("/about", H.About)

	router.GET("/login", H.Login)
	router.POST("/login", H.LoginPost)
	router.GET("/logout", H.Logout)

	admin := router.Group("/admin")
	admin.Use(middleware.CheckAuth)
	admin.GET("/add", H.Add)
	admin.POST("/add", H.AddPost)
	admin.GET("/update/:id", H.Update)
	admin.POST("/update", H.UpdatePost)
	admin.GET("/list", H.List)

	router.Run(":80")
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "views/shared/layout.html", "views/index.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("articles", "views/shared/layout.html", "views/articles.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("detail", "views/shared/layout.html", "views/detail.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("about", "views/shared/layout.html", "views/about.html", "views/shared/_header.html", "views/shared/_footer.html")
	r.AddFromFiles("reward", "views/shared/layout.html", "views/reward.html", "views/shared/_header.html", "views/shared/_footer.html")

	r.AddFromFiles("login", "views/auth/login.html")
	r.AddFromFiles("list", "views/shared/admin_layout.html", "views/admin/list.html", "views/shared/admin_header.html")
	r.AddFromFiles("add", "views/shared/admin_layout.html", "views/admin/add.html", "views/shared/admin_header.html")
	r.AddFromFiles("update", "views/shared/admin_layout.html", "views/admin/update.html", "views/shared/admin_header.html")
	return r
}
