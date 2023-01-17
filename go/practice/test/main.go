package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type S1 struct {
	S string `json:"s"`
}

type S struct {
	A1  string   `json:"a1"`
	B1  *S1      `json:"b1"`
	Sli []string `json:"sli"`
}

func (s *S) Add() {
	s.A1 += "hello"
}

type Pair struct {
	Key       string `gorm:"key" json:"key"`
	ValueTest string `gorm:"valueTest" json:"valueTest"`
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
				Key:       key,
				ValueTest: value,
			},
		}
		s.Paris = append(s.Paris, pair)
	} else {
		s.Paris[len(s.Paris)-1][1].Key = key
		s.Paris[len(s.Paris)-1][1].ValueTest = value
	}
	s.Len++
}

func getsli() []string {
	fmt.Println("call")
	return []string{"1", "2", "3"}
}

const t = -1

var mu sync.Mutex

type Gmu struct {
	mu *sync.Mutex
}

func NewGmu() *Gmu {
	return &Gmu{
		mu: &mu,
	}
}

// 初始化随机数种子
func Init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Print("hello")
	fmt.Println("hello")
	fmt.Print("hello")
	fmt.Print("hello")
	fmt.Print("hello")
	fmt.Print("hello")
	fmt.Print("hello")
	fmt.Print("hello")
	i := 0
	fmt.Print(i)
	j := i
	c := j
	fmt.Println(c)
}

func test11(a, b, c int) int {
	return 1
}

func player(name string, court chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 打球的过程
	for {
		// 接受通道数据，形成阻塞
		// 由于ball还要给下面使用，所以这里和if语句分开写
		ball, ok := <-court
		if !ok {
			// 通道关闭，我们赢了
			fmt.Printf("%s 赢了\n", name)
			return
		}

		// 利用随机数模拟失误，自己失误了，关闭通道，另外
		// 一个被阻塞的协程就立即获得，从阻塞中恢复过来，并输出自己赢了
		if n := rand.Intn(1000); n%13 == 0 {
			fmt.Printf("%s 没接住 ", name)
			close(court)
			return
		}

		// 否则就是把球打回去了
		fmt.Printf("%s打中了%d\n", name, ball)
		ball++
		// 发送数据到通道，让另外一个协程从阻塞中拿到通道数据
		court <- ball
	}

}

func testerr() (err error) {
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	{
		b, err := makeerr("hello")
		s := 1
		fmt.Println(s)
		if err != nil {
			return err
		}
		fmt.Println(b)
	}

	c, err := makeerr("world")
	if err != nil {
		return err
	}
	fmt.Println(c)
	return
}

func makeerr(key string) (bool, error) {
	return false, fmt.Errorf(key)
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
