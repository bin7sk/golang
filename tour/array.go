package main

import "fmt"

func main() {
	var array = [...]int{1, 2, 3, 4, 5}
	for i := range array {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()
	// slices
	slice := array[1:4]
	fmt.Println("slice:", slice)
	fmt.Println("length of slice:", len(slice))
	fmt.Println("capacity of slice:", cap(slice))
	slice = append(slice, 7)
	slice[0] = 0
	fmt.Println(array)
}
