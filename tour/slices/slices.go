package main

import "fmt"

func main() {
	arr := []string {
		"first",
		"second",
		"third",
		"fourth",
		"fiveth",
	}
	fmt.Println("Capacity: ", cap(arr))
	for i:=0; i<len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("slice:")
	slice := arr[1:4]
	fmt.Println("Capacity: ", cap(slice))
	fmt.Println("Length: ", len(slice))
	for i:=0; i<len(slice); i++ {
		fmt.Println(slice[i])
	}

}
