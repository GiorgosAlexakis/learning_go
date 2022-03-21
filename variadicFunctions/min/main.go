package main

import (
	"errors"
	"fmt"
)

func min(vals ...int) (int, error) {
	if vals == nil {
		return 0, errors.New("Provide at least one value")
	}
	min := vals[0]
	for _, val := range vals {
		if val <= min {
			min = val
		}
	}
	return min, nil
}

func main() {
	result, error := min(1, 2, 3, 4)
	fmt.Println(result)
	result, error = min()
	fmt.Println(result)
	fmt.Println(error)
}
