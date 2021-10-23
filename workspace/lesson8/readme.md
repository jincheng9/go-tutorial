# 函数

* 函数定义

  ```go
  func name([parameter list]) [return_types] {
    do sth
  }
  ```

  

  * 无参数

    ```go
    func name() int {
      do sth
    }
    ```

    

  * 无返回值

    ```go
    func name(a int) {
      do sth
    }
    ```

    

  * 返回1个值

    ```go
    func name(a int) int {
      do sth
    }
    ```

    

  * 返回多个值

    ```go
    func name(a int) (int, string) {
      do sth
    }
    func name(a b int) (int, string) {
      do sth
    }
    func name(a int, b string)(int, string) {
      do sth
    }
    func name(a, b int, c, d string) (int, string) {
      do sth
    }
    ```

    

* 函数参数

  * 值传递
  * 引用传递：指针

* 函数高级用法

  * 函数作为另一个函数的实参
  * 闭包
  * 方法