package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("main.c")
	defer file.Close()
	inputReader := bufio.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	for true {
		time.Sleep(time.Minute)
		inputString, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(inputString)
	}
}
