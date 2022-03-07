package quickSort

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	arr := []int{3, 4, 5, 667, 6, 3, 3, 1, 5, 345, 5, 4, 2, 435, 354}
	QuickSort(arr, 0, len(arr))
	fmt.Println(arr)
}
