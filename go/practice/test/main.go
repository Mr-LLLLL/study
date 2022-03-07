package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type S1 struct {
	S string
}

type S struct {
	A1  string `json:"a1"`
	B1  *S1    `json:"b1"`
	Sli []string
}

func (s *S) Add() {
	s.A1 += "hello"
}

type Pair struct {
	Key   string
	Value string
}

type Slice struct {
	Len   int
	Paris [][2]Pair
}

func (s *Slice) Append(key, value string) {
	isFull := s.Len&1 == 0
	if isFull {
		pair := [2]Pair{
			{
				Key:   key,
				Value: value,
			},
		}
		s.Paris = append(s.Paris, pair)
	} else {
		s.Paris[len(s.Paris)-1][1].Key = key
		s.Paris[len(s.Paris)-1][1].Value = value
	}
	s.Len++
}

func getsli() []string {
	fmt.Println("call")
	return []string{"1", "2", "3"}
}

const t = -1

func main() {

}

func Copy(dst, src interface{}) {
	b, err := json.Marshal(src)
	if err != nil {
		return
	}

	json.Unmarshal(b, dst)
}

func test() {
	arr := [1000000]int{1, 2, 3}
	for i := range arr[:] {
		_ = arr[i]
		time.Sleep(time.Second)
	}
}

func test1() {
	arr1 := [1000000]int{1, 1, 1}
	for i := range arr1 {
		_ = arr1[i]
		time.Sleep(time.Second)
	}
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(")finish", name)
			return
		default:
			fmt.Println("continue", name)
			time.Sleep(10 * time.Second)
		}
	}
}
