package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("pc[i/2]:   %08b\n", byte(pc[i/2]))
		//fmt.Printf("byte(i):   %08b\nbyte(1):   %08b\nbyte(i&1): %08b\n", byte(i), byte(1), byte(i&1))
		//fmt.Printf("pc[i]:     %08b\n\n", byte(pc[i]))
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func PopCountLoop(x uint64) (count int) {
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
func PopCountShift64(x uint64) (count int) {
	for i := 0; i < 64; i++ {
		count += int(x>>(i)) & 1
	}
	return count
}

func PopCountClearRightMostZero(x uint64) (count uint64) {
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func main() {
	fmt.Println(PopCount(2511112121212))
	fmt.Println(PopCountLoop(2511112121212))
	fmt.Println(PopCountShift64(2511112121212))
	fmt.Println(PopCountClearRightMostZero(2511112121212))
}
