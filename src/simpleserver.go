package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", echo)
	fmt.Println("Runnig server on 8000")
	http.ListenAndServe(":8000", nil)
}

func echo(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Request received")
	randNo := rand.Int()

	io.WriteString(rw, "Hello World!!!"+strconv.Itoa(randNo)+"\n")

}
