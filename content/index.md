go-water 是一款设计层面的框架，帮助 web 框架（gin，iris，beego，echo等）实现隔离友好，设计优美的系统，通过一系列接口、规范、约定，深度解耦系统。

### 安装
```
go get -u github.com/go-water/water
```

### 核心函数类型
```
type Endpoint func(ctx context.Context, req interface{}) (interface{}, error)
```
业务接口 Service 包含一个方法返回这个类型，见 Service 接口定义

### 介绍 Service 接口
```
type Service interface {
	Endpoint() Endpoint
	Name() string
	SetLogger(l *zap.Logger)
	GetRequest() interface{}
}
```
你所有的业务接口得都实现这个接口，这个是核心业务接口，如果自己定义对象new(IndexRequest)，业务服务可以不用提供 GetRequest 方法，由嵌套结构体提供，提供 GetRequest 方法可以减少内存分配，详情见低层代码

### 介绍内置的嵌套结构体 ServerBase
```
type ServerBase struct {
	l *zap.Logger
}

func (s *ServerBase) Name(srv interface{}) string
func (s *ServerBase) GetLogger() *zap.Logger
func (s *ServerBase) SetLogger(l *zap.Logger)
func (s *ServerBase) GetRequest() interface{}
```
这个结构体嵌套进业务结构体，使得业务结构体获得两个读写日志相关的方法，方法Name用来注入服务接口名，打印日志带上接口名更加友好

### 介绍 Handler 接口
```
type Handler interface {
	ServerWater(ctx context.Context, req interface{}) (interface{}, error)
	GetLogger() *zap.Logger
	GetRequest() interface{}
}
```
Handler 可以理解为接口 Service 的壳，它包装 Service，隐藏调用细节

### 如何创建一个具体的业务接口 Service (GetArticleService)
```
type GetArticleRequest struct {
	UrlID string `json:"url_id"`
}

type GetArticleService struct {
	*water.ServerBase
}

func (srv *GetArticleService) Handle(ctx context.Context, req *GetArticleRequest) (interface{}, error) {
	result, err := model.GetArticle(model.DbMap, req.UrlID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *GetArticleService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*GetArticleRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}

func (srv *GetArticleService) Name() string {
	return srv.ServerBase.Name(srv)
}
```
包含三个方法，其中两个方法是为了实现 Service 接口，还有一个方法 SetLogger，由于嵌套结构体已经实现，所以不用再实现，Handle 方法是获取数据层数据，或者其他业务数据

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
func (h *Handlers) GetArticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	req := new(service.GetArticleRequest)
	req.UrlID = id
	resp, err := h.getArticle.ServerWater(context.Background(), req)
	if err != nil {
		ctx.EndRequest()
	}

	ctx.ViewData("body", resp)
	ctx.View("detail.html")
}
```
把接口控制器函数写成 Handlers 方法，小写字母打头，避免字段与方法重名

### 日志处理
```
srv.GetLogger().Error(err.Error())
srv.GetLogger().Info("打印一条日志")
```
srv 就是业务实现 GetArticleService 的实例，在 GetArticleService 方法中，都可以打印日志。（这里封装了 zap 日志组件）

### 错误处理
```
type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}
```
每个业务服务接口，比如 GetArticleService 层，如果发生 error，低层会自动打印日志，日志里面会带上[GetArticleService]，以便区分，用户可以通过下面的 option 改写日志的方式，只需实现上面接口，然后在创建业务接口实现时改写行为。

### 配置 option
```
type ServerOption func(*Server)

// 改写低层错误处理
func ServerErrorHandler(errorHandler ErrorHandler) ServerOption
// 添加后置执行器
func ServerFinalizer(f ...ServerFinalizerFunc) ServerOption
// 配置日志Config
func ServerConfig(c *Config) ServerOption
```
结构体 Server 实现了 Handler 接口，配置 Server，其实是配置 Handler，上面代码来之样例仓库，经过简化处理，更加详细的文档和代码，请参考下面仓库和官方网站：

### 样例仓库
+ [https://github.com/go-water/go-water](https://github.com/go-water/go-water)

### 官方网站
+ [https://go-water.cn](https://go-water.cn)