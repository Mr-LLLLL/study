package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func condition(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("cond.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func iterator(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("iter.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, nil)
	t.Execute(w, daysOfWeek)
}

func argument(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("arg.html")
	t.Execute(w, "hello")
}

func include(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func pipe(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("pipe.html")
	t.Execute(w, "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/condition", condition)
	http.HandleFunc("/iterator", iterator)
	http.HandleFunc("/argument", argument)
	http.HandleFunc("/include", include)
	http.HandleFunc("/pipe", pipe)
	server.ListenAndServe()
}
