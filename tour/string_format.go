package main

import "fmt"

func main() {
	var integer int = 177

	// print value
	fmt.Printf("Typs: %T\n", integer)
	fmt.Printf("%-15s %20v\n", "value is:", integer)
	fmt.Printf("%-15s %20b\n", "binary:", integer)
	fmt.Println()



}
