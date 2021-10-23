# f循环控制

* for的4种用法

  * 用法1: 类似C++的for(int i=0; i<100; i++)

    ```go
    for init; condition; post {
      do sth
    }
    for ; condition; { // 类似用法2
      do sth
    }
    ```

  * 用法2：类似C++的while循环

    ```go
    for condition {
      do sth
    }
    ```

  * 用法3: 死循环，类似C++的for(;;)

    ```go
    for {
      do sth
    }
    ```

  * 用法4: For-each range循环, 类似python的 for k,v in dict.items()

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

  * break：跳出for循环或者switch控制逻辑

  * continue：结束当前循环，继续下一轮循环

* goto：类似C++里的goto

  * 语法

  ```go
  label: statement
  goto label
  ```

  * 代码示例

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

  

