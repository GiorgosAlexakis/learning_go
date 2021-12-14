package main

import "fmt"

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

func main() {
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("After: len(slice) =", len(slice))
}
