package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	s := "zello world"
	s1 := "hello, 中国"
	s2 := "'n'"
	fmt.Println(s, s1)
	fmt.Printf("%d\n", s[0])
	fmt.Printf("%d\n", s1[0])
	fmt.Println(s2)
	fmt.Println(len("世界"))
	fmt.Println(len("\xe4\xb8\x96\xe7\x95\x8c"))
	fmt.Println(len("\u4e16\u754c"))
	fmt.Println(len("\U00004e16\U0000754c"))
	fmt.Println(utf8.RuneCountInString("世界, hello"))
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Printf("%x\n", "中国你好")
	fmt.Printf("% x\n", 10000)
	fmt.Printf("%T\n", len("hello"))

	fmt.Println(basename("a/a/c.go"))
	fmt.Println(comma("1232456.78"))

	fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println([]int{1, 2, 3})
	fmt.Println(strconv.ParseInt("123", 5, 0))
	fmt.Printf("%f", math.Pi)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v %[1]d %d\n", time.Minute, len(strconv.Itoa(int(time.Minute))))

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	const (
		_   = 1 << (10 * iota)
		KiB = 1
		MiB = 2
		GiB = 3
		TiB
		PiB
		EiB
		ZiB
		YiB
	)
	fmt.Printf("Kib type is %T:\n", KiB)
	fmt.Println(YiB / ZiB)

	var x float32 = math.Pi
	var y float64 = math.Pi
	var xx float32 = float32(y)
	var z complex128 = math.Pi
	fmt.Println(x, y, z, xx)

	var f float64 = 212
	fmt.Printf("%T:\n", (f-32)*5/9)
	fmt.Printf("%T:\n", 5/9*(f-32))
	fmt.Printf("%T:\n", 5.0/9.0*(f-32))
	fmt.Printf("%T:\n", 123.9)
	fmt.Printf("%T:\n", math.Pi)
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	dot := strings.LastIndex(s, ".")
	var length int
	if dot == -1 {
		length = len(s)
	} else {
		length = dot
	}
	for i := length; i-3 > 0; i = i - 3 {
		s = s[:i-3] + "," + s[i-3:]
	}
	return s
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
