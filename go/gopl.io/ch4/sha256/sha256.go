package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%T, %[1]v\n", []byte("x"))

	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "october", 11: "Novermber", 12: "December"}
	Q2 := months[4:7]
	fmt.Println(Q2)
	summer := months[6:9]
	fmt.Printf("len = %d, cap = %d\n", len(summer), cap(summer))
	endlessSummer := summer[:5]
	fmt.Println(summer)
	fmt.Println(endlessSummer)

}
