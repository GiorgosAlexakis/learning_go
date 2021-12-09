package main

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))

	add_func := func(x, y int) int {
		return x + y
	}
	fmt.Println(add_func(1, 2))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	// defer function_whatever so function_whatever is executed at the end of the function
	z := 5
	zero(&z)
	fmt.Println(z)
	xPtr := new(int) //new takes a type as an argument, allocates enough memory to fit a value of that type,and returns a pointer to it
	one(xPtr)
	fmt.Println(*xPtr)
}
