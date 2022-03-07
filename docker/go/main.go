package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(writer, "Hello World, %s!, hostname:%s", request.URL.Path[1:], name)
}
