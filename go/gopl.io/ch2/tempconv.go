package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZereC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilngC       Celsius = 100
)

func main() {
    c := FToC(212.0)
    fmt.Println(c.String())
    fmt.Printf("%v\n", c)
    fmt.Printf("%s\n", c)
    fmt.Println(c)
    fmt.Printf("%g\n", c)
    fmt.Println(float64(c))

    var f Fahrenheit
    fmt.Println(c == 0)
    fmt.Println(f >= 0)
    // fmt.Printfln(c == f)     // compile error   two type not match
    fmt.Println(c == Celsius(f))

}

func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c * 9 / 5 + 32)
}

func FToC(f Fahrenheit) Celsius {
    return Celsius(( f - 32 ) * 5 / 9)
}

func (c Celsius) String() string {
    return fmt.Sprintf("%g℃", c)
}
