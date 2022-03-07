package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder
var post *FakePost

func TestMain(m *testing.M) {
	setUp()
	fmt.Println("testmain()===============")
	fmt.Println("testmain()---------------")
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	post = &FakePost{}
	mux.HandleFunc("/post/", handleRequest(post))
	writer = httptest.NewRecorder()
	fmt.Println("setup()===============")
}

func tearDown() {
	fmt.Println("tearDown()======================")
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/2", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 2 {
		t.Error("Cannot retrieve JSON psot")
	}
	fmt.Println("testhandleget()===============")
}

func TestHandlePut(t *testing.T) {
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(*post)
	fmt.Println("testhandleput()+++++++++++++++++++=")
}
