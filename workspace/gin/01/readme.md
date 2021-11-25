# Gin介绍

## 当前流行的Go语言web框架

以下是截止到2021.10.03，GitHub上开源的Go Web框架情况。目前Gin是遥遥领先。

| Project Name                                              | Stars | Forks | Open Issues | Description                                                  | Last Commit         |
| --------------------------------------------------------- | ----- | ----- | ----------- | ------------------------------------------------------------ | ------------------- |
| [gin](https://github.com/gin-gonic/gin)                   | 51894 | 5889  | 443         | Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. | 2021-09-30 02:04:28 |
| [beego](https://github.com/beego/beego)                   | 27032 | 5316  | 27          | beego is an open-source, high-performance web framework for the Go programming language. | 2021-09-18 15:08:26 |
| [kit](https://github.com/go-kit/kit)                      | 21360 | 2192  | 47          | A standard library for microservices.                        | 2021-09-28 15:01:29 |
| [echo](https://github.com/labstack/echo)                  | 20797 | 1841  | 65          | High performance, minimalist Go web framework                | 2021-09-26 15:56:43 |
| [fasthttp](https://github.com/valyala/fasthttp)           | 16135 | 1336  | 46          | Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http | 2021-10-01 11:38:31 |
| [fiber](https://github.com/gofiber/fiber)                 | 15590 | 789   | 37          | ⚡️ Express inspired web framework written in Go               | 2021-10-02 01:54:34 |
| [mux](https://github.com/gorilla/mux)                     | 15202 | 1413  | 24          | A powerful HTTP router and URL matcher for building Go web servers with 🦍 | 2021-09-14 12:12:19 |
| [kratos](https://github.com/go-kratos/kratos)             | 14913 | 3001  | 35          | A Go framework for microservices.                            | 2021-09-30 06:31:25 |
| [httprouter](https://github.com/julienschmidt/httprouter) | 13204 | 1275  | 63          | A high performance HTTP request router that scales well      | 2020-09-21 13:50:23 |
| [revel](https://github.com/revel/revel)                   | 12400 | 1402  | 103         | A high productivity, full-stack web framework for the Go language. | 2020-07-12 05:57:36 |
| [go-zero](https://github.com/zeromicro/go-zero)           | 11533 | 1372  | 34          | go-zero is a web and rpc framework written in Go. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. | 2021-10-02 10:16:59 |



## Web框架需要做什么

我们先思考下，一个完整的Web开发框架需要做哪些事情

| 组件            | 功能                                                         | 是否必须 |
| --------------- | ------------------------------------------------------------ | -------- |
| server          | 作为server，监听端口，接受请求                               | 是       |
| router          | 路由和分组路由，可以把请求路由到对应的处理函数               | 是       |
| middleware      | 支持中间件，对外部发过来的http请求经过中间件处理，再给到对应的处理函数。例如http请求的日志记录、请求鉴权(比如校验token)、CORS支持、CSRF校验等。 | 是       |
| template engine | 模板引擎，支持后端代码对html模板里的内容做渲染(render)，返回给前端渲染好的html | 否       |
| ORM             | 对象关系映射，可以把代码里的对象和关系数据库的表、字段做映射关联，通过操作对象来实现数据库的增删查改等操作。 | 否       |



## Gin有什么

Gin的主要作者是[Manu](https://github.com/manucorporat)，[Javier](https://github.com/javierprovecho)和[Bo-Yi Wu](https://github.com/appleboy)，2016年发布第一个版本，目前是最受欢迎的开源Go框架。

Gin除了支持上面表格里列的server、router、middleware和template之外，还支持

* Crash-free：崩溃恢复，Gin可以捕捉运行期处理http请求过程中的panic并且做recover操作，让服务一直可用。
* JSON validation：JSON验证。Gin可以解析和验证request里的JSON内容，比如字段必填等。当然开发人员也可以选择使用第三方的JSON validation工具，比如[beego validation](https://github.com/beego/beego/tree/develop/core/validation)。
* Error management：错误管理。Gin提供了一种简单的方式可以收集http request处理过程中的错误，最终中间件可以选择把这些错误写入到log文件、数据库或者发送到其它系统。
* Middleware Extendtable：可以自定义中间件。Gin除了自带的官方中间件之外，还支持用户自定义中间件，甚至可以把自己开发的中间件提交到[官方代码仓库](https://github.com/gin-gonic/contrib)里。

Gin本身不支持ORM，如果想在Gin框架里使用ORM，可以选择使用第三方的ORM，比如[gorm](https://github.com/go-gorm/gorm)。



# References

* https://gin-gonic.com/
* https://github.com/mingrammer/go-web-framework-stars

