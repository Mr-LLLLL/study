package main

import (
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
	}
}

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
	}
}

func printLetters2() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func print2() {
	printNumbers2()
	printLetters2()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
}
