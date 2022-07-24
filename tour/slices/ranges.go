package main

import "fmt"

var array = [] int {10,11,12,13,14,15,16,17,18,19,20,21,22}

func main() {
	for i, v:= range array {
		fmt.Printf("array[%d] = %d\n", i, v)
	}
}
