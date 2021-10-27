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

func printBook2(book *Book) {
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
    
    book3 := Book{1, "expert", "go"}
    bookPtr := &book3
    printBook2(bookPtr)
}