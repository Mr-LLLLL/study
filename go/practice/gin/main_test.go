package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:12345/test", nil)
	if nil != err {
		t.Fatal(err)
	}

	r, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
}
