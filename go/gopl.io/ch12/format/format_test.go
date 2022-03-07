package format

import (
	"fmt"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Second
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))
}
