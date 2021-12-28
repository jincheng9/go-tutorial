# 循环控制

## for的4种用法

* 用法1: 类似C++的`for(int i=0; i<100; i++)`

  ```go
  for init; condition; post {
    do sth
  }
  
  for init; condition; {
    do sth
  }
  
  for ; condition; { // 类似下面的用法2
    do sth
  }
  ```

* 用法2：类似C++的while循环

  ```go
  for condition {
    do sth
  }
  ```

* 用法3: 死循环，类似C++的`for(;;)`

  ```go
  for {
    do sth
  }
  ```

* 用法4: For-each range循环, 类似python的` for k,v in dict.items()`

  可以对slice，map，数组和字符串等数据类型进行For-each迭代循环

  ```go
  for key, value := range map1 { // 遍历map
    do sth
  }
  for index, value := range list { // 遍历数组
    do sth
  }
  for index, character := range str { // 遍历字符串
    do sth
  }
  ```

* break：跳出当前for循环或者switch控制逻辑

* continue：结束当前循环，继续下一轮for循环

## goto：

类似C++里的goto

* 语法

```go
label: statement
goto label
```

* 代码示例1

```go
package main

import "fmt"

func main() {
	LOOP: 
		println("Enter your age:")
		var age int
		_, err := fmt.Scan(&age) // 接受控制台输入
		if err != nil {
			println("error:", err)
			goto LOOP
		}
		if age < 18 {
			println("You are not eligible to vote!")
			goto LOOP
		} else {
			println("You are eligible to vote!")
		}
		println("all finish")
}
```



## break和label结合

break和label结合使用，可以跳出二重或者多重for循环。

例1：`break A`直接跳出整个外层for循环，所以下面的例子只执行`i=0, j=0`这一次循环。

```go
package main
import "fmt"

// 最终输出 0 0 Hello, 世界
func main() {
A:
	for i := 0; i < 2; i++ {
		for j := 0; i < 2; j++ {
			print(i, " ", j, " ")
			break A
		}

	}
	fmt.Println("Hello, 世界")
}
```

例2：下面的例子，break只能跳出位于里层的for循环，会执行`i=0, j=0`和`i=1, j=0`这2次循环。

```go
package main

import "fmt"
// 输出 0 0 1 0 Hello, 世界
func main() {
	for i := 0; i < 2; i++ {
		for j := 0; i < 2; j++ {
			print(i, " ", j, " ")
			break
		}

	}
	fmt.Println("Hello, 世界")
}
```



