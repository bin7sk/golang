//implemented 'selection sorting' algorithm
package main

import "fmt"

func main() {
	// array with random numbers
	// sort from min to max
	array_to_sort := []int{
		3, 7, 1,0, 140, 20, 220, 12, 10, 11, 46, 4,
		57, 93, 20, 11, 7, 4, 5, 6, 1, 0, 2, 40, 160, 
		30, 105, 205, 103, 277, 311, 88, 77, 66, 49}
	length := len(array_to_sort)
	for i := 0; i < length - 1; i++ {
		min_index := i
		for j := i; j < length; j++ {
			if array_to_sort[j] < array_to_sort[min_index] {
				min_index = j
			}
		}	
		if (min_index != i) {
			tmp := array_to_sort[i]
			array_to_sort[i] = array_to_sort[min_index]
			array_to_sort[min_index] = tmp 
		}
	}
	
	//print array
	for i:=0; i<length; i++ {
		fmt.Printf("%d ", array_to_sort[i])
	}
	fmt.Println()

}
