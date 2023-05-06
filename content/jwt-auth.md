### 创建token
```
SetAuthToken(uniqueUser, privateKeyPath string, expire time.Duration) (tokenString string, err error)
```
输入参数
+ uniqueUser，这个是用户标识符，可以用一个唯一的数字id，也可以是其他用户账号，但必须全局唯一
+ privateKeyPath，秘钥的路径
+ expire，超时时间

输出参数
+ tokenString，生成的token
+ err，错误

### 验证token
```
ParseAndValid(req *http.Request, publicKeyPath string) (uniqueUser, signature string, err error)
```
输入参数
+ req，go低层类型，每个web框架都包含了这个对象
+ publicKeyPath，公钥的路径

输出参数
+ uniqueUser，这个是用户标识符，对应于登陆输入参数uniqueUser
+ signature，这个签名只包含jwt的最后一段字符串
+ err，错误

### 用法简介
在用户登陆时，提交账号，密码，如果正确就调用 SetAuthToken 创建 token，将 token 发送给客户端，客户端保留 token，以便下次请求使用

在验证阶段，用户将 token 信息包含在 header 中，通常使用关键字“Authorization”，如：Authorization: Bearer token，Bearer 和 token 之间有一个空格

### 私钥，公钥生成
```
# 生成私钥
openssl genrsa -out rsa_private.key
# 生成公钥
openssl rsa -in rsa_private.key -pubout -out rsa_public.key
```
openssl命令一般linux系统都内置，可以直接用