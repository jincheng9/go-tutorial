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