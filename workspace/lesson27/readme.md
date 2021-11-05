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

**注意**：这里说的是同一个package，不是同一个文件。同一个package下，如果有多个源程序文件是申明的该package，那这些源程序文件里的变量、函数、结构体等，即使不是首字母大写，也可以互相跨文件直接调用，不用额外import。

package的使用分为3类情况：

* 使用Go标准库自带的package，比如fmt。
* 使用go get获取到的第三方package
* 使用工程本地的package



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

  在Go 1.11之前，还没有Go Modules，如果想import一些自己开发的package，Go应用必须建在$GOPATH/src路径下。

### 开启GO111MODULE

* 先从$GOROOT/src里找。

* 如果$GOROOT/src找不到，再看当前项目有没有go.mod文件，有的话就从go.mod文件里指定的模块所在路径往下找。如果没有go.mod文件，那就直接提示package xxx is not in GOROOT。

  

官方推荐使用Go Modules，从Go1.16版本开始，默认开启GO111MODULES环境变量为on模式。



## 在Module中使用package





## 注意事项

* package目录名和package目录下的Go源程序文件开头声明的包名可以不一样，不过一般还是写成一样，避免出错。

## References

* https://www.callicoder.com/golang-packages/
* https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/
* https://maelvls.dev/go111module-everywhere/#go111module-with-go-116

