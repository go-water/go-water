go-water 的配置参考 go-kit 与 go-micro 设计，使用 option 设置定义，所有定义的 option 都是作为 Service 服务的参数，见末尾使用方式

### option 函数类型
```
type ServerOption func(*Server)
```

### 自定义错误处理
```
type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

func ServerErrorHandler(errorHandler ErrorHandler) ServerOption {
	return func(s *Server) { s.errorHandler = errorHandler }
}
```
用例
```
errorHandler := 一个实现 ErrorHandler 接口的结构体实例
option := water.ServerErrorHandler(errorHandler)
```
可以自定义一个日志处理方式，比如写入磁盘

### 后置执行器
```
type ServerFinalizerFunc func(ctx context.Context, err error)

func ServerFinalizer(f ...ServerFinalizerFunc) ServerOption {
	return func(s *Server) { s.finalizer = append(s.finalizer, f...) }
}
```
用例
```
finalizerFunc := 一个 ServerFinalizerFunc 类型函数
option := water.ServerFinalizer(finalizerFunc)
```
请求执行尾部需要执行的函数，可以理解为后置执行器

### 日志配置
```
func ServerConfig(c *Config) ServerOption {
	return func(s *Server) { s.c = c }
}
```
用例
```
conf := &water.Config{Encoding: "console", Level: zap.InfoLevel}
option := water.ServerConfig(conf)
return &Handlers{
	index:       water.NewHandler(&service.IndexService{ServerBase: &water.ServerBase{}}, option),
}
```
说明：每个 Handler 要实例一个 water.ServerBase，嵌套到 Service 服务里，不要共用一个实例。