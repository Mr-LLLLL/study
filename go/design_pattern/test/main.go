package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := t{
		a: 1,
		B: 2,
	}
	var c t
	b, _ := json.Marshal(a)
	json.Unmarshal(b, &c)
	fmt.Println(c)
}

type t struct {
	a int
	B int
}

