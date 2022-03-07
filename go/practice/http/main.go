package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	Post()
}

type Body map[string]interface{}

func Post() {
	m := Body{
		"trash": Body{
			"isTrash": false,
		},
		"idOrName": "1478999361219203072",
	}

	body, _ := json.Marshal(m)

	req, err := http.NewRequest("POST", "https://mk-test.dustess.com/mall/goods/v1/list", strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", "qw.2a8baaae-6e99-11ec-8ac5-9e661db32914")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))
}
