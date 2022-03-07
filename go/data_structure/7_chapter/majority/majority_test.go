package majority

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	arr := []int{1, 2, 1, 1, 1, 1, 1, 1, 1}
	arr1 := []int{1, 2, 1, 2, 1, 2}

	fmt.Println(Majority(arr), Majority(arr1))
}
