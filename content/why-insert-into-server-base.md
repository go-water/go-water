用例
```go
return &Handlers{
	index:       water.NewHandler(&service.IndexService{ServerBase: &water.ServerBase{}}, option),
}
```
嵌套一个结构体，主要是代码重用，减少重复代码，同时还可以注入自定义 option 到框架层，更容易达到想要的效果