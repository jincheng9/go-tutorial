package main

import (
	"encoding/json"
	"fmt"
)

type Status uint32

const (
	StatusOpen Status = iota
	StatusClosed
	StatusUnknown
)

type Request struct {
	ID        int    `json:"Id"`
	Timestamp int    `json:"Timestamp"`
	Status    Status `json:"Status"`
}

func main() {
	// marshal, convert struct to json string
	req := Request{ID: 1, Timestamp: 10}
	str, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))

	// unmarshal, convert json string to struct
	bs := []byte(`{"ID":1, "Timestamp": 12}`)
	req2 := Request{}
	fmt.Println(req2)
	if err := json.Unmarshal(bs, &req2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(req2)
}
