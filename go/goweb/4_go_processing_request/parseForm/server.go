package main

import (
	"fmt"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	// fmt.Fprintln(w, r.PostForm)

	// r.ParseMultipartForm(1024)
	// fmt.Fprintln(w, r.MultipartForm)

	// fmt.Fprintln(w, r.FormValue("hello"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	log.Fatal(server.ListenAndServe())
}
