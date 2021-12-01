# Gin源码结构解析

阅读本文之前，可以先读[上一篇文章](../01)，对Web框架的核心组件有个了解。

## 源代码的目录结构

以v1.7.4版本为例，Gin的源码目录结构如下图所示

```go
+---.github
|       ISSUE_TEMPLATE.md
|       PULL_REQUEST_TEMPLATE.md
|       
+---binding
|       binding.go
|       binding_msgpack_test.go
|       binding_nomsgpack.go
|       binding_test.go
|       default_validator.go
|       default_validator_test.go
|       form.go
|       form_mapping.go
|       form_mapping_benchmark_test.go
|       form_mapping_test.go
|       header.go
|       json.go
|       json_test.go
|       msgpack.go
|       msgpack_test.go
|       multipart_form_mapping.go
|       multipart_form_mapping_test.go
|       protobuf.go
|       query.go
|       uri.go
|       validate_test.go
|       xml.go
|       xml_test.go
|       yaml.go
|       yaml_test.go
|       
+---examples
|       README.md
|       
+---ginS
|       gins.go
|       README.md
|       
+---internal
|   +---bytesconv
|   |       bytesconv.go
|   |       bytesconv_test.go
|   |       
|   |---json
|           json.go
|           jsoniter.go
|           
+---render
|       data.go
|       html.go
|       json.go
|       msgpack.go
|       protobuf.go
|       reader.go
|       reader_test.go
|       redirect.go
|       render.go
|       render_msgpack_test.go
|       render_test.go
|       text.go
|       xml.go
|       yaml.go
|       
|---testdata
|    +---certificate
|    |       cert.pem
|    |       key.pem
|    |       
|    +---protoexample
|    |       test.pb.go
|    |       test.proto
|    |       
|    |---template
|            hello.tmpl
|            raw.tmpl
|   .gitignore
|   .travis.yml
|   auth.go
|   AUTHORS.md
|   auth_test.go
|   BENCHMARKS.md
|   benchmarks_test.go
|   CHANGELOG.md
|   codecov.yml
|   CODE_OF_CONDUCT.md
|   context.go
|   context_appengine.go
|   context_test.go
|   CONTRIBUTING.md
|   debug.go
|   debug_test.go
|   deprecated.go
|   deprecated_test.go
|   doc.go
|   errors.go
|   errors_1.13_test.go
|   errors_test.go
|   fs.go
|   gin.go
|   gin_integration_test.go
|   gin_test.go
|   githubapi_test.go
|   go.mod
|   go.sum
|   LICENSE
|   logger.go
|   logger_test.go
|   Makefile
|   middleware_test.go
|   mode.go
|   mode_test.go
|   path.go
|   path_test.go
|   README.md
|   recovery.go
|   recovery_test.go
|   response_writer.go
|   response_writer_test.go
|   result.txt
|   routergroup.go
|   routergroup_test.go
|   routes_test.go
|   test_helpers.go
|   tree.go
|   tree_test.go
|   utils.go
|   utils_test.go
|   version.go
```

使用[cloc](http://cloc.sourceforge.net/)工具对源码做一个扫描，总共代码行数不到1.2W。

```go
http://cloc.sourceforge.net v 1.64  T=0.91 s (97.2 files/s, 18596.7 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              84           2537           2360          11796
make                             1              8              0             63
YAML                             2              8              0             47
Protocol Buffers                 1              2              0             10
-------------------------------------------------------------------------------
SUM:                            88           2555           2360          11916
-------------------------------------------------------------------------------
```



## 目录结构说明

| 目录     | 作用                                                         |
| -------- | ------------------------------------------------------------ |
| 根目录   | 整个Web框架的核心都在根目录下面，包括server, router, middleware。 |
| binding  | 把http request的数据转成Go里自定义的结构体变量，可以不用自己去逐个解析http request里的参数，减少开发工作量。 |
| examples | 使用Gin的一些代码示例，目前该目录只有一个README.md，具体的例子在[Gin Examples](https://github.com/gin-gonic/examples)这个单独的Repo里。 |
| ginS     | 依赖sync.Once实现了一个单例版本的gin.Engine，并且对gin.go里的方法做了简单的二次封装。平时做测试可以使用下，不建议生产使用，可以忽略。关于ginS在 GitHub我开过一个issue做讨论，感兴趣的可以通过后面的[ginS链接](https://github.com/gin-gonic/gin/issues/2957)进行查看。 |
| internal | Gin内部用的函数。在json这个子package里引用了标准库的encoding/json和第三方的[json-iterator](https://github.com/json-iterator/go)，在bytesconv这个子package包里实现了零内存分配版本的string和byte切片的互相转换。 |
| render   | 支持将XML、JSON、YAML、ProtoBuf以及HTML数据做渲染，返回给前端可识别的响应格式。 |
| testdata | 一些测试数据。                                               |



## 根目录源码说明

按照源码重要程度排序

| 源码                                | 作用                                                         |
| ----------------------------------- | ------------------------------------------------------------ |
| gin.go                              | gin.Engine的定义和方法实现，Gin的入口程序                    |
| context.go<br>context_appengine.go  | gin.Context的定义和方法实现                                  |
| response_write.go                   | 响应数据的处理                                               |
| routergroup.go<br>tree.go           | 路由的具体实现，其中tree.go来源于[httprouter](https://github.com/julienschmidt/httprouter/blob/master/tree.go)，基于radix tree，不使用反射 |
| auth.go<br>recovery.go<br>logger.go | 中间件，auth.go是http鉴权中间件，recovery.go是程序panic捕获中间件，logger.go是日志中间件 |
| mode.go                             | Gin服务启动模式的初始化管理，支持debug, release和test这3种模式 |
| debug.go                            | mode.go开启debug模式时，打印一些调试信息                     |
| fs.go                               | 和静态文件访问相关的辅助代码，主要是给routergroup.go的Static方法用 |
| path.go                             | 对url路径做处理的辅助函数                                    |
| utils.go                            | 一些辅助函数                                                 |
| version.go                          | 记录当前Gin框架的版本信息                                    |
| deprecated.go                       | 未来要被删掉的方法，目前只有一个BindWith方法。不应该在程序里使用deprecated.go里的方法。 |
| doc.go                              | 目前没有实际用途，可忽略                                     |
| xx__test.go                         | 测试程序                                                     |

在后续的文章里，会对Gin的每个源代码文件做详细解析。
