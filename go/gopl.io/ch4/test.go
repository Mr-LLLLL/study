package main

import (
	"fmt"

	"gopl.io/ch4/wheel"
)

func main() {
	var w wheel.Wheel
	w.Y = 10
	fmt.Printf("%#v\n", w)
	fmt.Printf("%22s world", "helloworld")
}
