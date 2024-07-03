package main

import (
	"fmt"
)

type str struct {
	id string
}

func main() {
	fmt.Printf("%[1]s%[1]s", "sdf")
}

func hello() {
	fmt.Println()
}
