package main

import (
	"html/template"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html", "tmpl1.html")
	t.ExecuteTemplate(w, "tmpl1.html", "Hello World!\n")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	log.Fatal(server.ListenAndServe())
}
