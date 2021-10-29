# 接口interface

* 定义：接口是一种抽象的类型，是一组method的集合，里头只有method方法，没有数据成员。当两个或两个以上的类型都有相同的处理方法时才需要用到接口。先定义接口，然后多个struct类型去实现接口里的方法，就可以通过接口变量去调用struct类型里实现的方法。

  比如动物都会叫唤，那可以先定义一个名为动物的接口，接口里有叫唤方法speak，然后猫和狗这2个struct类型去实现各自的speak方法。

* 语法：

  ```go
  // 定义接口
  type interface_name interface {
    method_name1([参数列表]) [返回值列表]
    method_name2([参数列表]) [返回值列表]
    method_nameN([参数列表]) [返回值列表]
  }
  
  // 定义结构体类型
  type struct_name struct {
      data_member1 data_type
      data_member2 data_type
      data_memberN data_type
  }
  
  // 实现接口interface_name里的方法method_name1
  func(struct_var struct_name) method_name1([参数列表])[返回值列表] {
      /*具体方法实现*/
  }
  
  // 实现接口interface_name里的方法method_name2
  func(struct_var struct_name) method_name2([参数列表])[返回值列表] {
      /*具体方法实现*/
  }
  
  /* 实现接口interface_name里的方法method_name3
  注意：下面用了指针接受者。函数可以使用值接受者或者指针接受者，上面的method_name1和method_name1使用的是值接受者。
  如果用了指针接受者，那给interface变量赋值的时候要传引用
  */
  func(struct_var *struct_name) method_name3([参数列表])[返回值列表] {
      /*具体方法实现*/
  }
  
  ```

  

* 示例：

  ```go
  package main
  
  import "fmt"
  
  // all animals can speak
  type Animal interface {
      speak()
  }
  
  // cat
  type Cat struct {
      name string
      age int
  }
  
  func(cat Cat) speak() {
      fmt.Println("cat miaomiaomiao")
  }
  
  // dog
  type Dog struct {
      name string
      age int
  }
  
  func(dog *Dog) speak() {
      fmt.Println("dog wangwangwang")
  }
  
  
  func main() {
      var animal Animal = Cat{"gaffe", 1}
      animal.speak() // cat miaomiaomiao
      
      /*
      因为Dog的speak方法用的是指针接受者，因此给interface赋值的时候，要赋指针
      */
      animal = &Dog{"caiquan", 2}
      animal.speak() // dog wangwangwang
  }
  ```

* 多个类型可以实现同一个interface：多个类型都有共同的方法(行为)。比如上面示例里的猫和狗都会叫唤，猫和狗就是2个类型，叫唤就是speak方法。

* 一个类型可以实现多个interface：一个类型可能有好几个

* 空interface

  * 如果空interface作为函数参数，可以接受任何类型的实参
  * 如果空interface作为变量，可以把任何类型的变量赋值给空interface

* **注意事项**

