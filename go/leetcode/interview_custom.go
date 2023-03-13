package leetcode

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 1、构造一个具有1000万个随机10-110范围内的整数的切片，并编写一个求和函数，求切片中所有元素的和，并计算求和需要的时间。要求充分利用Golang的并发特性。
// NOTE: use hash to sum
// 求和时间复杂度和空间复杂度分别为O(1), O(1)
type T struct {
	data [101]int
	num  int
}

func NewT() *T {
	t := &T{
		num: 1000 * 10000,
	}
	t.genData()

	return t
}

func (t *T) genData() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(101)
	for i := 0; i <= 100; i++ {
		index := i
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case val := <-ch:
					t.data[index] = val
					return
				}
			}

		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()

		randBase := t.num / 101 * 2
		remaind := t.num
		rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < 100 && remaind > 0; i++ {
			r := rand.Intn(randBase) + 1
			ch <- r
			remaind -= r
		}
		if remaind != 0 {
			ch <- remaind
		}
	}()
	wg.Wait()
}

func (t *T) GetSli() []int {
	wg := sync.WaitGroup{}
	res := make([]int, t.num)
	ch := make(chan int, 100)
	wg.Add(202)
	for i := 0; i <= 100; i++ {
		index := i
		go func() {
			defer wg.Done()

			for i := t.data[index]; i > 0; i-- {
				ch <- index + 10
			}
		}()

		go func() {
			defer wg.Done()

			base := t.num/101 + 1
			for i := 0; i < base; i++ {
				if i+base*index >= t.num {
					return
				}
				res[i+base*index] = <-ch
			}
		}()
	}
	wg.Wait()

	return res
}

func (t *T) GetSum() int {
	sum := 0
	for i, v := range t.data {
		sum += (i + 10) * v
	}
	return sum
}

// 2、使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728。
func Print() string {
	var (
		ping        = make(chan struct{}, 1)
		pong        = make(chan struct{}, 1)
		ctx, cancel = context.WithCancel(context.Background())
		ch          = make(chan string)
		res         = ""
	)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 26; {
			<-pong
			ch <- fmt.Sprintf("%c%c", 'A'+i, 'A'+i+1)
			i += 2
			ping <- struct{}{}
		}
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; ; {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case <-ping:
				ch <- strconv.Itoa(i) + strconv.Itoa(i+1)
				i += 2
				pong <- struct{}{}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range ch {
			res += v
		}
	}()

	ping <- struct{}{}
	wg.Wait()

	return res
}
