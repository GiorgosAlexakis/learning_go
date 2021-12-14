package main

import "fmt"

func Insert(slice []int, index, value int) []int {
	// Grow the slice by one element.
	slice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}
func main() {
	slice := make([]int, 10, 15)
	fmt.Printf("len: %d, cap:%d\n", len(slice), cap(slice))
	newSlice := make([]int, len(slice), 2*cap(slice))
	/* Not advisable to use a loop to copy data
	for i := range slice {
		newSlice[i] = slice[i]
	}
	slice = newSlice
	*/
	// It is better to use copy instead
	// copy returns an integer value,the number of elements it copied
	copy(newSlice, slice)
	fmt.Printf("len: %d, cap:%d\n", len(slice), cap(slice))
	slice_1 := make([]int, 10, 20) // Note capacity > length: room to add element.
	for i := range slice {
		slice_1[i] = i
	}
	fmt.Println(slice_1)
	slice_1 = Insert(slice_1, 5, 99)
	fmt.Println(slice_1)
}
