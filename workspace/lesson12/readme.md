# 结构体

* 定义

  * 语法

    ```go
    type struct_type struct {
        member_name1 data_type1
        member_name2 data_type2
        member_name3, member_name4 data_type3
    }
    
    struct_var := struct_type{value1, value2, value3, value4}
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

    

* 访问结构体内的成员使用 点**.**   **结构体变量.成员**

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

  

* 结构体指针

  * 语法。注意结构体指针访问结构体里的成员，也是用点**.**，这个和C++用->不一样

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

    

* 方法

* 大写字母

