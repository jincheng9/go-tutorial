# 条件语句

* 布尔表达式可以不加括号

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

* switch

  * 每一个case分支都是唯一的，从上往下逐一判断，知道匹配为止。如果某些case分支条件重复了，编译会报错

  * 默认情况下每个case分支最后自带break效果，匹配成功就不会执行其它case。

    如果需要执行后面的case，可以使用fallthrough。使用 fallthrough 会强制执行紧接着的下一个 case 语句，fallthrough 不会判断紧接着的下一条 case 的表达式结果是否为 true。

  * 方法1

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

  * 方法2

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