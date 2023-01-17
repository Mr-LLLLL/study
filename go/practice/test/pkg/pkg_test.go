package pkg

import (
	"fmt"
	"testing"
	"time"
)

// Test_test function
func Test_test(t *testing.T) {
	fmt.Println("sdf")
	fmt.Println("sdf")
	fmt.Println("sdf")
	fmt.Println("sdf")
	time.Sleep(1 * time.Second)
}

func Test_test2(t *testing.T) {
	fmt.Print("sdf")
}

func TestLogPath(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogPath()
		})
	}
}
