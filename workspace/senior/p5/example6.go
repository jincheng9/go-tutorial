// example6.go
// 请求命令：curl -X POST -H "content-type: application/json" http://localhost:4000/user -d '{"name":"John", "age":111}'
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	// data 匿名结构体变量，用来接收http请求发送过来的json数据
	data := struct{
		Name string `json:"name"`
		Age int	`json:"age"`
	}{}
	// 把json数据反序列化到data变量里
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(data)
	fmt.Fprint(w, "Hello!")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user", AddUser)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
