package main

import "fmt"

func swap(x, y string) (string, string) {
	return y,x
}

func main() {
	a, b := swap("first arg", "second arg")
	fmt.Println(a)
	fmt.Println(b)
}
