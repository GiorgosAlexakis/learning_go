package main

import "fmt"

// func arrayOfIntegers creates an array of 5 integers
// set the value of the last element to 100
// print the array
func arrayOfIntegers() {
	var a [5]int
	a[4] = 100
	fmt.Println("a:", a)
}

// func arrayOfStrings creates an array of 5 empty strings
// set y[0] to "Hello"
// set y[1] to "World"
// set y[2] to "!"
// print the array from index 0 to index 3
func arrayOfStrings() {
	var y [5]string
	y[0] = "Hello"
	y[1] = "World"
	y[2] = "!"
	fmt.Println("y:", y[0:3])
}

//func calculateAverage takes an array of integers
//and returns the average of those integers
func calculateAverage(scores []int) float64 {
	var sum int
	for _, score := range scores {
		sum += score
	}
	return float64(sum) / float64(len(scores))
}

func main() {
	arrayOfIntegers()
	arrayOfStrings()
	var average float64 = calculateAverage([]int{98, 93, 77, 82, 83})
	fmt.Println("Average of test scores:[98, 93, 77, 82, 83] is", average)
}
