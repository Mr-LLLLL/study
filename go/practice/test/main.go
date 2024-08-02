package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type str struct {
	Msg string `json:"traceID"`
}

type str1 struct {
	Msg string `json:"msg"`
}

func main() {
	f1, _ := os.OpenFile("/home/ubuntu/workspace/xmirror/client-command-line-tool/src/test1", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o664)
	f2, _ := os.OpenFile("/home/ubuntu/workspace/xmirror/client-command-line-tool/src/test", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o664)

	scanner1 := bufio.NewScanner(f1)
	m := make(map[string]bool)
	for scanner1.Scan() {
		tmp := new(str1)
		json.Unmarshal(scanner1.Bytes(), tmp)

		parts := strings.Split(tmp.Msg, " ")
		m[parts[1]] = true
	}

	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		tmp := new(str)
		json.Unmarshal(scanner2.Bytes(), tmp)

		if _, ok := m[tmp.Msg]; !ok {
			fmt.Println(tmp.Msg)
		}
	}
}
