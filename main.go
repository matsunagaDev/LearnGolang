package main

import "fmt"
func incrementGenerator() (func() int) {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
 	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
  f := 1.11
	fmt.Println(f)
	fmt.Println(int(f))

	m := map[string]int{
		"Mike": 20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}
