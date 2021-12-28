# 条件语句

## If

布尔表达式可以不加括号

* if/else

  ```go
  if expression {
    do sth1
  } else {
    do sth2
  }
  ```

* if/else if/else

  ```go
  if expression1 {
    do sth1
  } else if expression2 {
    do sth2
  } else {
    do sth3
  }
  ```

* if/else嵌套

  ```go
  if expression1 {
    if expression11 {
      do sth11
    } else {
      do sth12
    }
  } else if expression2 {
    do sth2
  } else {
    do sth3
  }
  ```

## switch

* 每一个case分支都是唯一的，从上往下逐一判断，直到匹配为止。如果某些case分支条件重复了，编译会报错

* 默认情况下每个case分支最后自带break效果，匹配成功就不会执行其它case。

  如果需要执行后面的case，可以使用fallthrough。

  使用 fallthrough 会强制执行紧接着的下一个 case 语句，不过fallthrough 不会去分析紧接着的下一条 case 的表达式结果是否满足条件，而是直接执行case里的语句块。

  ```go
  // Foo prints and returns n.
  func Foo(n int) int {
      fmt.Println(n)
      return n
  }
  
  func main() {
      switch Foo(2) {
      case Foo(1), Foo(2), Foo(3):
          fmt.Println("First case")
          fallthrough
      case Foo(4):
          fmt.Println("Second case")
      }
  }
  ```

  比如上面的例子，执行结果如下，并不会去执行`fallthrough`的下一个case分支里的表达式`Foo(4)`

  ```markdown
  2
  1
  2
  First case
  Second case
  ```

* switch使用方法1

  ```go
  switch variable {
    case value1:
      do sth1
    case value2:
      do sth2
    case value3, value4: // 可以匹配多个值，只要一个满足条件即可
      do sth34
    case value5:
      do sth5
    default:
    	do sth
  }
  ```

* switch使用方法2

  ```go
  switch os := runtime.GOOS; os {
  	case "darwin":
  		fmt.Println("OS X.")
  	case "linux":
  		fmt.Println("Linux.")
  	default:
  		// freebsd, openbsd,
  		// plan9, windows...
  		fmt.Printf("%s.\n", os)
  }
  
  // 上面的写法和这个等价
  os := runtime.GOOS
  switch os {
  	case "darwin":
  		fmt.Println("OS X.")
  	case "linux":
  		fmt.Println("Linux.")
  	default:
  		// freebsd, openbsd,
  		// plan9, windows...
  		fmt.Printf("%s.\n", os)
  }
  ```
  
* switch使用方法3。case分支的每个condition结果必须是一个bool值，要么为true，要么为false

  ```go
  switch {
    case condition1:
    	do sth1
    case condition2:
    	do sth2
    case condition3, condition4:
    	do sth34
    default:
    	do sth
  }
  ```

* switch使用方法4。只适用于`interface`的类型判断，而且`{`要和`switch`在同一行，`{`前面不能有分号`;`

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	var i interface{} = 10
  	switch t := i.(type) {
  	case bool:
  		fmt.Println("I'm a bool")
  	case int:
  		fmt.Println("I'm an int")
  	default:
  		fmt.Printf("Don't know type %T\n", t)
  	}
  }
  ```

## References

* https://yourbasic.org/golang/switch-statement/