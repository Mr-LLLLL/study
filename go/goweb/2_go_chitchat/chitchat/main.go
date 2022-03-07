package main

import (
	"goweb/2_go_chitchat/chitchat/route"
	"goweb/2_go_chitchat/chitchat/utilities"
	"log"
	"net/http"
	"time"
)

func main() {
	utilities.PrintAppInfo()

	mux := route.MakeRouter()

	// starting up the server
	server := &http.Server{
		Addr:           utilities.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(utilities.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(utilities.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatalln(server.ListenAndServe())
}
