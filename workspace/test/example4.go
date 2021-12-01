package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("xtapp.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix | log.LUTC)
	log.SetPrefix("INFO:")
	log.Println(12)
}