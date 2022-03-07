package route

import (
	"goweb/2_go_chitchat/chitchat/utilities"
	"net/http"
)

func MakeRouter() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	// handle static assetst
	files := http.FileServer(http.Dir(utilities.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	/*
	* all route patterns matched here
	* route handler functions defined in other files
	 */

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccout)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)
	return
}
