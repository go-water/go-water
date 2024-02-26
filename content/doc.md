简介：go-water 是一款设计层面的 web 框架（像 gin，iris，beego，echo 一样，追求卓越）。 我们使命：更好的业务隔离，更好的系统设计，通过一系列接口、规范、约定、中间件，深度解耦业务系统。

### 安装
```
go get -u github.com/go-water/water
```
go-water 除了实现 web 框架必要的组件以外，还实现了业务设计，每个业务接口（即一个数据请求/api/getData，或者一个视图页面/homePage）抽象成一个 Handler 接口，和一个 Service 接口。Handler 在应用层创建，Service 在业务层创建。

### 介绍 Service 接口
```
type Service interface {
	Name(srv Service) string
	SetLogger(l *slog.Logger)
}
```
你所有的业务接口得都实现这个接口，这个是核心业务接口，同时业务服务还包含一个嵌套的ServerBase，自动获得它的方法

### 介绍内置的嵌套结构体 ServerBase
```
type ServerBase struct {
	l *slog.Logger
}

func (s *ServerBase) Name(srv Service) string
func (s *ServerBase) GetLogger() *slog.Logger
func (s *ServerBase) SetLogger(l *slog.Logger)
```
这个结构体嵌套进业务结构体，丰富业务服务的功能，简化代码，使得业务结构体获得两个读写日志相关的方法，方法Name用来注入服务接口名，打印日志带上接口名更加友好

### 介绍 Handler 接口
```
type Handler interface {
	ServerWater(ctx context.Context, req any) (any, error)
	GetLogger() *slog.Logger
}
```
Handler 可以理解为接口 Service 的代理接口，它包装 Service，隐藏调用细节

### 如何创建一个具体的业务接口 Service (GetArticleService)，经过简化，保留核心代码
```
type GetArticleRequest struct {
	UrlID string `json:"url_id"`
}

type GetArticleService struct {
	*water.ServerBase
}

func (srv *GetArticleService) Handle(ctx context.Context, req *GetArticleRequest) (interface{}, error) {
	article := new(Article)
	return article, nil
}
```
这个结构体由于嵌套结构体，所以它实现了接口 Service，所以不用再实现，Handle 方法是获取数据层数据

### 创建一个 Handler，并归入 Handlers 结构体
```
type Handlers struct {
	getArticle  water.Handler
}

func NewService() *Handlers {
	return &Handlers{
		getArticle:  water.NewHandler(&GetArticleService{ServerBase: &water.ServerBase{}}),
	}
}
```
每个业务接口可以理解为一个 Handler，每个业务接口实现可以理解为一个 Service，创建 Handler 就是将 Service 接口实现作为参数传递给 water.NewHandler，嵌套一个 ServerBase 可以重复减少代码量

### 控制器层调用
```
func (h *Handlers) GetArticle(ctx *water.Context) {
	id := ctx.Param("id")
	req := new(service.GetArticleRequest)
	req.UrlID = id
	resp, err := h.getArticle.ServerWater(ctx, req)
	if err != nil {
		h.getArticle.GetLogger().Error(err.Error())
		return
	}

	title := ""
	if article, ok := resp.(*model.Article); ok {
		title = article.Title
	}

	ctx.HTML(http.StatusOK, "detail", water.H{"body": resp, "title": title})
}
```
把接口控制器函数写成 Handlers 方法，小写字母打头，避免字段与方法重名

### 日志处理
```
srv.GetLogger().Error(err.Error())
srv.GetLogger().Info("打印一条日志")
```
srv 就是业务实现 GetArticleService 的实例，在 GetArticleService 方法中，都可以打印日志。（这里返回 slog 日志实例）

### 配置 option
```
type ServerOption func(h *handler)

type ServerFinalizerFunc func(ctx context.Context, err error)

func ServerFinalizer(f ...ServerFinalizerFunc) ServerOption {
	return func(h *handler) { h.finalizer = append(h.finalizer, f...) }
}

func ServerLimiter(interval time.Duration, b int) ServerOption {
	return func(h *handler) {
		h.limit = rate.NewLimiter(rate.Every(interval), b)
	}
}

func ServerBreaker(breaker *gobreaker.CircuitBreaker) ServerOption {
	return func(h *handler) {
		h.breaker = breaker
	}
}
```
结构体 handler 实现了 Handler 接口，配置 handler，其实是配置 Handler

### JWT 集成
```
// 创建 token
func SetAuthToken(uniqueUser, issuer, privateKeyPath string, expire time.Duration) (tokenString string, err error)

// 验证 token，兼容 http,ws
func ParseFromRequest(req *http.Request, publicKeyPath string) (uniqueUser, issuer, signature string, err error)
```

### 限流，通过 option 将限流的中间件用上
```
func NewService() *Handlers {
	option := water.ServerLimiter(time.Minute, 100)
	return &Handlers{
		getArticle:  water.NewHandler(&service.GetArticleService{ServerBase: &water.ServerBase{}}, option),
	}
}
```

### 注意
仓库代码已更新到 v0.8.4

### 架构源码
+ [https://github.com/go-water/water](https://github.com/go-water/water)

### 样例仓库
+ [https://github.com/go-water/go-water](https://github.com/go-water/go-water)

### 官方网站
+ [https://go-water.cn](https://go-water.cn)