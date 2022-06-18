package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	// The fmt.Printf verbs %b,%d,%u,%x are often more
	// convenient than Format functions, especially if
	// we want to include information besides the number
	s := fmt.Sprintf("x=%b", x) // "x=1111011"
	fmt.Println(s)
	x, err := strconv.Atoi("123")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(x + 1)
	y_dec, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(y_dec)
	y_bin, err := strconv.ParseInt("010101", 2, 64) // result is always int64
	if err != nil {
		log.Println(err)
	}
	fmt.Println(y_bin)
}
