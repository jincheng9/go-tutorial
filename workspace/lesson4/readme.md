# 常量

* 常量定义的时候必须赋值，定义后值不能被修改

* 常量(包括全局常量和局部常量)可以定义后不使用，局部变量定义后必须使用，否则编译报错

* 常量可以用来定义枚举

* iota，特殊常量，可以理解为const语句块里的行索引，值从0开始

* 常量的定义方法

  * 方法1

    ```go
    const a int = 10
    const b bool = false
    ```

  * 方法2

    ```go
    const a = 10
    const b = false
    ```

  * 多个常量同时定义

    ```go
    const a, b int = 1, 2
    ```

  * iota，特殊常量，可以理解为每个独立的const语句块里的行索引

    ```go
    const a int = iota // the value of a is 0
    const b = iota // the value of b is still 0
    ```

  * 定义枚举方法1

    ```go
    const (
      unknown = 0
      male = 1
      female = 2
    )
    ```

  * 定义枚举方法2

    ```go
    const (
      unknown = iota // the value of unknown is 0
      male // the value of male is 1
      female // the value of female is 2
    )
    const (
      c1 = iota // the value of c1 is 0
      c2 = iota // the value of c2 is 1
      c3 = iota // the value of c3 is 2
    )
    ```

  * 注意事项

    * iota的值是const语句块里的行索引，行索引从0开始
    * const语句块里，如果常量没赋值，那它的值和上面的保持一样，比如下面的例子里class2=0, class6="abc"
    * 某个常量赋值为iota后，紧随其后的常量如果没赋值，那后面常量的值是自动+1，比如下面的例子里，class3的值是iota，该行的行索引是2，所以class3=2， class4常量紧随其后没有赋值，那class4=class3+1=3

    ```go
    const (
    	class1 = 0
    	class2 // class2 = 0
    	class3 = iota  //iota is 2, so class3 = 2
    	class4 // class4 = 3
    	class5 = "abc" 
    	class6 // class6 = "abc"
    	class7 = iota // class7 is 6
    )
    ```

    

