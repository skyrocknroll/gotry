package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fzzy/radix/redis"
)

func main() {

	http.HandleFunc("/", health)
	http.HandleFunc("/health", health)

	fmt.Println("Server listening on 7379")
	http.ListenAndServe(":7379", nil)
}
func health(rw http.ResponseWriter, req *http.Request) {
	r, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
	// If there is timeout error then send 504 error
	if err != nil {

		rw.WriteHeader(http.StatusGatewayTimeout)
		io.WriteString(rw, err.Error())
		fmt.Println(err.Error())
		return
	}
	defer r.Close()
	data, err := r.Cmd("INFO").Str()
	// Send 500 error if INFO command fails
	if err != nil {
		// io.WriteString(rw, err.Error())
		// log.Fatal(err)

		rw.WriteHeader(http.StatusInternalServerError)
		io.WriteString(rw, err.Error())
		fmt.Println(err.Error())
	} else {
		io.WriteString(rw, data)
	}
}
