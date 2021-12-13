// gin.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
	Info struct{
		Age int `json:"age"`
	} `json:"info"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header["Content-Type"])
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	//data, err := ioutil.ReadAll(r.Body)

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
