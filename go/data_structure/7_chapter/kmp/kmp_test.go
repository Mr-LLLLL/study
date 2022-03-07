package algorithm

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	patter := "hello"
	strs := []string{
		"helloworld",
		"worldhello",
		"1234",
		"123hello",
		"ksjdfkafksakfaskfjksfskjf",
	}
	for _, s := range strs {
		fmt.Println(Match([]byte(s), []byte(patter)))
	}
}
