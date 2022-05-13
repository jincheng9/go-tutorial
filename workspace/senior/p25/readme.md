# Go 1.18：工作区模式workspace mode

## 背景

Go 1.18除了引入泛型(generics)、模糊测试(Fuzzing)之外，另外一个重大功能是引入了工作区模式(workspace mode)。

工作区模式对**本地同时开发多个有依赖的Module**的场景非常有用。

举个例子，我们现在有2个Go module项目处于开发阶段，其中一个是`example.com/main`，另外一个是`example.com/util`。其中`example.com/main`这个module需要使用`example.com/util`这个module里的函数。

我们来看看Go 1.18版本前后如何处理这种场景。



## Go 1.18之前怎么做

在Go 1.18之前，对于上述场景有2个处理方案：

### 方案1：被依赖的模块及时提交代码到代码仓库

这个方案很好理解，既然`example.com/main`这个module依赖了`example.com/util`这个module，那为了`example.com/main`能使用到`example.com/util`的最新代码，需要做2个事情

* 本地开发过程中，如果`example.com/util`有修改，都马上提交代码到代码仓库，然后打tag
* 紧接着`example.com/main`更新依赖的`example.com/util`的版本号(tag)，可以使用`go get -u`命令。

这种方案比较繁琐，每次`example.com/util`有修改，都要提交代码，否则`example.com/main`这个module就无法使用到最新的`example.com/util`。



### 方案2：go.mod里使用replace指令

为了解决方案1的痛点，于是有了方案2：在go.mod里使用replace指令。

通过replace指令，我们可以直接使用`example.com/util`这个module的本地最新代码，而不用把`example.com/util`的代码提交到代码仓库。

为了方便大家理解，我们直接上代码。代码目录结构如下：

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

main目录下的`main.go`代码如下：

```go
//main.go
package main

import (
	"fmt"

	"example.com/util"
)

func main() {
	result := util.Add(1, 2)
	fmt.Println(result)
}
```

main目录下的`go.mod`内容如下：

```markdown
module example.com/main

go 1.16

require example.com/util v1.0.0

replace example.com/util => ../util
```

util目录下的`util.go`代码如下：

```go
// util.go
package util

func Add(a int, b int) int {
	return a + b
}
```

util目录下的`go.mod`内容如下：

```markdown
module example.com/util

go 1.16
```

这里最核心的是`example.com/main`这个module的`go.mod`，最后一行使用了replace指令。

```markdown
module example.com/main

go 1.16

require example.com/util v1.0.0

replace example.com/util => ../util
```

通过replace指令，使用go命令编译代码的时候，会找到本地的util目录，这样`example.com/main`就可以使用到本地最新的`example.com/util`代码。进入main目录，运行代码，结果如下所示：

```bash
$ cd main
$ go run main.go
3
```

但是这种方案也有个问题，我们在提交`example.com/main`这个module的代码到代码仓库时，需要删除最后的replace指令，否则其他开发者下载后会编译报错，因为他们本地可能没有util目录，或者util目录的路径和你的不一样。

代码开源地址：[Go 1.18之前使用replace指令](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p25/module)



## Go 1.18工作区模式

为了解决方案2的痛点，在Go 1.18里新增了工作区模式(workspace mode)。

该模式下不再需要在`go.mod`里使用replace指令，而是新增一个`go.work`文件。

话不多说，直接上代码。代码目录结构如下：

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

main目录下的`main.go`代码如下：

```go
//main.go
package main

import (
	"fmt"

	"example.com/util"
)

func main() {
	result := util.Add(1, 2)
	fmt.Println(result)
}
```

main目录下的`go.mod`内容如下：没有了方案2里最后一行的replace指令

```markdown
module example.com/main

go 1.16

require example.com/util v1.0.0
```

util目录下的`util.go`代码如下：

```go
// util.go
package util

func Add(a int, b int) int {
	return a + b
}
```

util目录下的`go.mod`内容如下：

```markdown
module example.com/util

go 1.16
```

`go.work`内容如下：

```markdown
go 1.18

use (
	./main
	./util
)
```

在workspace目录下执行如下命令即可自动生成`go.work`

```bash
$ go1.18beta1 work init main util
```

`go1.18beta1 work init`后面跟的`main`和`util`都是Module对应的目录。

如果go命令执行的当前目录或者父目录有`go.work`文件，或者通过`GOWORK`环境变量指定了`go.work`的路径，那go命令就会进入工作区模式。在工作区模式下，go就可以通过`go.work`下的module路径找到并使用本地的module代码。

在main目录或者workspace目录，都可以运行`main.go`，结果如下所示：

```bash
$ go1.18beta1 run main/main.go 
3
$ cd main/
$ go1.18beta1 run main.go
3
```

这种模式下，我们对`example.com/main` 没有任何本地侵入性修改，不用像方案2那样，提交代码前还需要更新`go.mod`文件。`example.com/main`里的内容都可以直接提交到代码仓库。

代码开源地址：[Go 1.18工作区模式](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p25/workspace)

**注意**：`go.work`不需要提交到代码仓库中，仅限于本地开发使用。



## 总结

为了解决多个有依赖的Module本地同时开发的问题，Go 1.18引入了工作区模式。

工作区模式是对已有的Go Module开发模式的优化，关于工作区模式的更多细节可以参考本文最后的References。



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
* go mod tidy不能感知go.work问题：https://github.com/golang/go/issues/50750