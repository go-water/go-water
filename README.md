# go-water
[go-water/water](https://github.com/go-water/water) 框架用例，详细文档和相关资料

### 星星增长趋势
[![Stargazers over time](https://starchart.cc/go-water/go-water.svg)](https://starchart.cc/go-water/go-water)

### 技术栈
+ gin
+ zap
+ gorp
+ viper
+ go-water
+ markdown

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
手动创建文件，修改数据库配置

### 安装步骤
+ 根据 model/article 创建数据表
+ 手动添加几条数据供 demo 使用
+ 请确保 80 端口没有被别的服务占用，然后在浏览器中输入：http://localhost

### 官方文档
+ https://iissy.com/go-water