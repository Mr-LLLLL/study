package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	defer func() {
		fmt.Println("==================================================")
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
