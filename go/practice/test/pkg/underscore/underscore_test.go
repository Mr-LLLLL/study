package underscore

import (
	"bytes"
	"fmt"
	"testing"

	un "github.com/tobyhede/go-underscore"
)

func Test_Each(t *testing.T) {
	// var Each func(func(value interface{}, i interface{}), interface{})

	var buffer bytes.Buffer

	fn := func(s, i interface{}) {
		buffer.WriteString(s.(string))
	}

	s := []string{"a", "b", "c", "d", "e"}
	un.Each(fn, s)

	fmt.Println(buffer.String())
}

func Test_EachInt(t *testing.T) {
	var sum int

	fn := func(v, i int) {
		sum += v
	}

	i := []int{1, 2, 3, 4, 5}
	un.EachInt(fn, i)

	fmt.Println(sum)
}

func Test_EachStringInt(t *testing.T) {
	var sum int
	fn := func(v int, k string) {
		sum += v
	}
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	un.EachStringInt(fn, m)

	fmt.Println(sum)
}

func Test_Map(t *testing.T) {
	s := []string{"a", "b", "c", "d"}

	fn := func(s interface{}) interface{} {
		return s.(string) + "!"
	}

	m := un.Map(fn, s)
	fmt.Println(m)
}

func Test_MapString(t *testing.T) {
	s := []string{"a", "b", "c", "d"}

	var SMap func([]string, func(string) string) []string
	un.MakeMap(&SMap)

	fn := func(s string) string {
		return s + "!"
	}

	m := un.MapString(fn, s)
	fmt.Println(m)
}

func Test_Partitino(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(i interface{}) bool {
		return (i.(int) % 2) == 1
	}

	odd, even := un.Partition(fn, s)

	fmt.Println(odd)  //[1, 3, 5, 7, 9]
	fmt.Println(even) //[2, 4, 6, 8, 10]
}

func Test_PartitionInt(t *testing.T) {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(v, i int) {
	}

	odd, even := un.PartitionInt(fn, s)

	fmt.Println(odd)  //[1, 3, 5, 7, 9]
	fmt.Println(even) //[2, 4, 6, 8, 10]
}
