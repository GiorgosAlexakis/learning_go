package main

import "fmt"



func main() {

	x := make([]float64, 5)
	y := make([]float64, 5, 10)
	fmt.Println(x, y)
	slice1 := []int{2, 3, 5, 7, 11, 13}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

}
