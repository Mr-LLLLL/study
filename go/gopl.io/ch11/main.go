package main

import "fmt"

func main() {
	test()
	test1()
	test2()
	test11()
	fmt.Println("hellow ol")
}

func test() {
	test1()
}

func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {

}

func test11() {
	test2()
}
