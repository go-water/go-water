cd ../go-water.cn
if exist server del server
cd ../go-water
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go build -o ../go-water.cn/server main.go
pause