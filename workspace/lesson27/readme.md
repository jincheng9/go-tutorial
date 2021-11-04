# 包package

## 定义

package本质上就是一个目录，目录里包含有一个或者多个Go源程序文件，或者package。也就是说package里面还可以嵌套包含子package。

每个Go源文件都属于一个package，在源文件开头指定package名称

```go
package package_name
```

package的代码结构大致如下：



package里的变量、函数，结构体，方法要被其它程序引用，需要使用大写字母。



## 在Module中使用package



## package嵌套



## 注意事项

* package目录名和package目录下的Go源程序文件开头声明的包名可以不一样，不过一般还是写成一样，避免出错
* 

## References

* https://www.callicoder.com/golang-packages/

* https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/