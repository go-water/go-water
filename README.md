# go-water
[github.com/go-water/water](https://github.com/go-water/water) 框架用例，详细文档和相关资料

### 星星增长趋势
[![Stargazers over time](https://starchart.cc/go-water/go-water.svg)](https://starchart.cc/go-water/go-water)

### config 目录下缺一个 mysql.yaml
```
mysql:
  go-water:
    db: "root:123456@tcp(192.168.1.1:3306)/go-water?parseTime=true&loc=Local&charset=utf8"
    max_idle_conns: 30
    max_open_conns: 100
    conn_max_lifetime: 100
    is_show_log: false
```
手动创建文件，并修改修改数据库账号，密码

### 安装步骤
+ 数据库初始化文件 init.sql，可以一键创建数据库 go-water
+ 请确保 80 端口没有被别的服务占用，然后在浏览器中输入：http://localhost

### 官方网站
+ https://jitask.com