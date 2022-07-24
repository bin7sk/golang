package main

import "fmt"

func adder() func(int) int {
	sum  := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

func main() {
	a := adder()
	a(10)
	a(5)
	fmt.Println(a(0))
	b := adder()
	b(1)
	b(1)
	fmt.Println(b(1))
}
