# 包package

## 定义

package本质上就是一个目录，目录里包含有一个或者多个Go源程序文件，或者package。也就是说package里面还可以嵌套包含子package。

每个Go源文件都属于一个package，在源文件开头指定package名称

```go
package package_name
```

package的代码结构示例如下：

![image-20211104181754164](./img/package_structure.jpg)

package里的变量、函数、结构体、方法等如果要被本package外的程序引用，需要在命名的时候首字母大写。

如果首字母小写，那就只能在同一个package里面被使用。

**注意**：这里说的是同一个package，不是同一个文件。同一个package下，如果有多个源程序文件是声明的该package，那这些源程序文件里的变量、函数、结构体等，即使不是首字母大写，也可以互相跨文件直接调用，不用额外import。

package的使用分为3类情况：

* 使用Go标准库自带的package，比如fmt。
* 使用go get获取到的第三方package
* 使用工程本地的package



## import语法示例

### 普通

```go
import (
	"fmt"                           // 标准库
	"sync/atomic"                   // 标准库sync的atomic子package
	"package1"                      // 自开发的package
	"package2/package21"            // 自开发package，嵌套子package
	"package2/package22"            // 自开发package，嵌套子package
	"package3/package31/package311" // 自开发package，多重嵌套
)
```

使用import路径里面定义的**package名称**来访问package里的方法，结构体等，而不是路径名称。

举个例子，假设上面import的路径package2/package21这个目录下的Go源程序文件开头声明的package名称是realpackage，那访问这个package里的方法，结构体等要用realpackage.xxx来访问，而不是用package21.xxx来访问。

一句话总结：**import的是路径，访问用package名称**。最佳实践就是让两者保持一致。

### 别名

```go
import (
    "fmt"
    newName "package2/package21"
)
```

可以用别名newName来访问package里的成员，newName.xxx。这个在包名很长或者包名有重复的时候可以用到。

### 点操作

```go
import (
    "fmt"
    . "package2/package21"
)
```

** . **可以让后面的package里的成员注册到当前包的上下文，这样就可以直接调用成员名，不需要加包前缀了。

比如以前要用package21.Hello()来调用package21这个包里的函数Hello，用了点操作后，可以直接调用函数Hello()，前面不用跟package名称。



### 下划线

```go
import (
    "fmt"
    _ "package2/package21"
)
```

下划线** _ **的效果：只会执行包里各个源程序文件的init方法，没法调用包里的成员。



## Go如何寻找import的package

在代码里import某个package的时候，Go是如何去寻找对应的package呢？这个和Go环境变量GO111MODULE有关系。GO111MODULE的值可以通过如下命令查到

```go
go env | grep GO111MODULE
```

on表示开启，off表示关闭。GO111MODULE是从Go 1.11开始引入，在随后的Go版本中Go Modules的行为有一些变化，具体可以参考[GO111MODULE and Go Modules](https://maelvls.dev/go111module-everywhere/#go111module-with-go-116)。

下面以Go1.16及以上版本详细讲下GO111MODULE关闭和开启的情况下，Go是如何寻找import的package的。

### 关闭GO111MODULE

* 先从$GOROOT/src里找。$GOROOT是Go的安装路径，$GOROOT/src是Go标准库存放的路径，比如fmt, strings等package都存放在$GOROOT/src里。$GOROOT的路径可以通过下面的命令查看到：

  ```go
  go env | grep ROOT // linux or mac
  go env | findstr ROOT // windows
  ```

* 如果从$GOROOT/src找不到，再从$GOPATH/src里找。$GOPATH是安装Go后就会有的一个环境变量，Linux和Mac的默认路径是/Users/用户名/go，WIndows默认路径是C:/Users/用户名/go

  ```go
  go env | grep PATH // linux or mac
  go env | findstr PATH // windows
  ```

  在Go 1.11之前，还没有Go Modules，如果想import一些自己开发的package，被import的package必须建在$GOPATH/src路径下。一般而言，一个工程项目一定会有自己写的若干个package，因此这也导致工程项目本身也通常建在了$GOPATH/src路径下。

### 开启GO111MODULE

Go 1.11开始，有了Go Modules，工程项目可以建在任何地方，代码在import某个package的时候，会按照如下顺序寻找package：

* 先从$GOROOT/src里找。

* 如果$GOROOT/src找不到，再看当前项目有没有go.mod文件，有的话就从go.mod文件里指定的模块所在路径往下找。如果没有go.mod文件，那就直接提示package xxx is not in GOROOT。

  

官方推荐使用Go Modules，从Go1.16版本开始，GO111MODULES环境变量默认开启为on模式。



## 使用示例

### 不开启GO111MODULES时import package

1. 项目建在$GOPATH/src下面
2. import package的时候路径从$GOPATH/src往下找

使用说明参考[gopath package demo](./gopath/)



### 开启GO111MODULES时import package

1. 项目可以建在任何地方

2. 在项目所在根目录创建go.mod文件

   ```go
   go mod init module_name
   ```

3. import项目里的package的时候指定go.mod文件里的模块名称

使用说明参考[module package demo](./module)



## init函数

init函数没有参数，没有返回值。

* 每个package里可以有多个init函数

* 每个源程序文件里也可以有多个init函数

* init函数不能被显示调用，在main()函数执行之前，自动被调用
* 同一个pacakge里的init函数调用顺序不确定
* 不同package的init函数，根据package import的依赖关系来决定调用顺序，比如package A里import了package B，那package B的init()函数就会比package A的init函数先调用。



## 注意事项

* package目录名和package目录下的Go源程序文件开头声明的包名可以不一样，不过一般还是写成一样，避免出错。

* 禁止循环导入package。

  

## References

* https://www.callicoder.com/golang-packages/
* https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/
* https://maelvls.dev/go111module-everywhere/#go111module-with-go-116

