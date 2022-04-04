# 结构体

## 定义

* 语法

  ```go
  type struct_type struct {
      member_name1 data_type1
      member_name2 data_type2
      member_name3, member_name4 data_type3
  }
  
  // 方式1：必须给结构体里每个成员赋值，如果只给部分成员赋值会编译报错
  struct_var := struct_type{value1, value2, value3, value4}
  // 方式2：可以给部分或者全部成员赋值，没有赋值的成员的值是成员所属类型的零值
  struct_var2 := struct_type{member_name1:value1, member_name2:value2}
  ```

  

* 示例

  ```go
  type Book struct {
      id int
      title string
      author string
  }
  
  book1 := Book{1, "go tutorial", "jincheng9"}
  book2 := Book{id:2, title:"day day up", author:"unknown"}
  ```

  

## 访问结构体内的成员

访问结构体内的成员使用点`.`   ，格式为：**结构体变量`.`成员**

```go
package main

import "fmt"

type Book struct {
    id int
    author string
    title string
}

var book Book

func printBook(book Book) {
    fmt.Println("id:", book.id)
    fmt.Println("author:", book.author)
    fmt.Println("title:", book.title)
}

func main() {
    book1 := Book{1, "go tutorial", "jincheng9"}
    book2 := Book{id:2, title:"day day up", author:"unknown"}
    printBook(book1)
    printBook(book2)
    
    book.id = 10
    book.author = "test"
    book.title = "test"
    printBook(book)
}
```



## 结构体指针

* 语法。**注意**:结构体指针访问结构体里的成员，也是用点`.`，这个和C++用`->`不一样

  ```go
  var struct_pointer *struct_type // 指针struct_pointer指向结构体struct_type
  struct_var := struct_type{} // 结构体变量
  struct_pointer = &struct_var // 指针赋值
  ```

  

* 示例

  ```go
  package main
  
  import "fmt"
  
  type Book struct {
      id int
      author string
      title string
  }
  
  func printBook(book *Book) {
      fmt.Println("id:", book.id)
      fmt.Println("author:", book.author)
      fmt.Println("title:", book.title)
  }
  
  func main() {
      book := Book{1, "expert", "go"}
      bookPtr := &book
      printBook(bookPtr)
  }
  ```

  

## 方法

* Go没有C++的class概念，但是可以对struct结构体类型定义方法，结构体对象调用该方法，来达到类似效果

  ```go
  package main
  
  import "fmt"
  
  type Book struct {
      id int
      author string
      title string
  }
  
  
  func (book Book) printBook() {
      fmt.Printf("id:%d, author:%s, title:%s\n", book.id, book.author, book.title)
  }
  
  func (book *Book) changeTitle1() {
      book.title = "new title1"
  }
  
  // 这个无法改变调用该方法的结构体变量里的成员的值
  func (book Book) changeTitle2() {
      book.title = "new title2"
  }
  
  func main() {
      book := Book{1, "expert", "go"}
      book.printBook()
      
      book.changeTitle1() // 会修改变量book里的成员title的值
      book.printBook()
      
      book.changeTitle2() // 不会对book的值有任何影响
      book.printBook()
      
  }
  ```


## 可见性

* 结构体标识符和结构体的成员标识符可见性
  * 如果结构体要被其它package使用，那结构体的标识符或者说结构体的名称首字母要大写
  * 如果结构体的成员要被其它package使用，那结构体和结构体的成员标识符首字母都要大写，否则只能在当前包里使用