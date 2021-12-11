# 设置GOPROXY地址

GOPROXY控制GO Module下载的来源。

Go默认的GOPROXY地址是https://proxy.golang.org,direct

通过如下命令可以获取到对应的值

```go
go env | grep GOPROXY // linux or mac
go env | findstr GOPROXY // windows
```

由于某些原因，在国内访问默认GOPROXY地址会失败，国内有一些好用的代理地址可以替换

七牛: https://goproxy.cn

阿里云：https://mirrors.aliyun.com/goproxy/

官方：https://goproxy.io



替换默认GOPROXY命令，以七牛的地址为例：

```go
go env -w GOPROXY=https://goproxy.cn,direct
```



