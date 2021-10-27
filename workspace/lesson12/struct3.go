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