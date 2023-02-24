用例
```
func (srv *IndexService) Handle(ctx context.Context, req *IndexRequest) (interface{}, error) {
	srv.GetLogger().Info("流水日志记录")
	// 你的业务代码，省略
	
	return nil, nil
}
```
所有服务调用 error，在内部方法 ServerWater 都会写日志，并记录服务名，用户如果需要写流水日志可以如上调用