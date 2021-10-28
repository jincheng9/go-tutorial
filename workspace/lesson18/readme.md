# 接口interface

* 定义：接口是一种抽象的类型，是一组method的集合。当两个或以上的类型都有相同的处理方法时才定义接口，先定义接口，然后多个struct类型去实现接口里的方法，就可以通过接口变量去调用struct类型里实现的方法。

  比如动物都会叫唤，那可以先定义一个名为动物的接口，接口里有叫唤方法speak，然后猫和狗这2个struct类型去实现各自的speak方法。

* 语法：

  ```go
  type interface_name interface {
    method_name1([参数列表]) [返回值列表]
    method_name2([参数列表]) [返回值列表]
    method_name3([参数列表]) [返回值列表]
  }
  ```

  

* 示例：

* 为什么要用interface

* 一个类型可以实现多个interface

* 多个类型可以实现同一个interface

* 空interface

* **注意事项**

