# Go1.18：工作区模式workspace mode

## 背景

Go 1.18除了引入泛型(generics)、模糊测试(Fuzzing)之外，另外一个重大功能是引入了工作区模式(workspace mode)。

工作区模式对**本地同时开发多个有依赖的Module**的场景非常有用。

举个例子，我们现在有2个Go module项目正在开发中，其中一个是`example.com/main`，另外一个是`example.com/util`。其中`example.com/main`这个module需要使用`example.com/util`这个module里的函数。

我们来看看在Go 1.18版本前后如何处理这种场景。



## Go 1.18之前怎么做

在Go 1.18之前，如果某个Go Module要引用本地的其它Module时，需要在`go.mod`里使用replace指令

* 每次都提交
*  replace指令

```markdown
module
|
|------main
|        |---main.go
|        |---go.mod        
|------util
|        |---util.go
|        |---go.mod
```

代码开源地址：[Go 1.18之前使用replace指令](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p25/module)



## Go 1.18工作区模式



```markdown
workspace
|------go.work
|
|------main
|        |---main.go
|        |---go.mod        
|------util
|        |---util.go
|        |---go.mod
```



代码开源地址：[Go 1.18工作区模式](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p25/workspace)



## 总结





## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* Proposal提案：https://go.googlesource.com/proposal/+/master/design/45713-workspace.md

* workspace官方教程：https://go.dev/doc/tutorial/workspaces
* workspace语法：https://go.dev/ref/mod#go-work-file-replace
* go work命令手册：https://pkg.go.dev/cmd/go@master#hdr-Workspace_maintenance
* go 1.18 release notes: https://tip.golang.org/doc/go1.18
* Go如何引用本地Module：https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson27#%E5%BC%80%E5%90%AFgo111modules%E6%97%B6import%E6%9C%AC%E5%9C%B0%E7%9A%84module
* polarisxu: https://polarisxu.studygolang.com/posts/go/dynamic/go1.18-workspace/