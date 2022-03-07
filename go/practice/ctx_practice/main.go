package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// withCancel()
	// withDeadline()
	// withValue()
	var b bool
	fmt.Println(b)
}

func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	earNum := eatHamburger(ctx)
	for n := range earNum {
		if n >= 10 {
			break
		}
	}
	cancel()

	fmt.Println("counting ... ")
	time.Sleep(time.Second * 2)
}

func eatHamburger(ctx context.Context) <-chan int {
	c := make(chan int)
	n := 0
	t := 0
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("cost time %d second eat %d hamburger\n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n > 10 {
					n = 10
				}
				t++
				fmt.Printf("i am eat %d hamburger\n", n)
			}
		}
	}()

	return c
}

func withDeadline() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	eatHamburger1(ctx)
	defer cancel()
}

func eatHamburger1(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡\n", n)
		}
		time.Sleep(time.Second)
	}
}

func withValue() {
	ctx := context.WithValue(context.Background(), "trace_id", "88888888")
	// 携带session到后面的程序中去
	ctx = context.WithValue(ctx, "session", 1)

	process(ctx)
}

func process(ctx context.Context) {
	session, ok := ctx.Value("session").(int)
	fmt.Println(ok)
	if !ok {
		fmt.Println("something wrong")
		return
	}

	if session != 1 {
		fmt.Println("session 未通过")
		return
	}

	traceID := ctx.Value("trace_id").(string)
	fmt.Println("traceID:", traceID, "-session:", session)
}
