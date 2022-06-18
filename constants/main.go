package main

import (
	"fmt"
	"math"
	"time"
)

const pi = math.Pi

const IPv4Len = 4

// since constant values are known to the compiler, constant expressions may appear
// in types:
// func ParseIPv4(s string) IP {
// 		var p [IPv4Len]byte
//		...
// }

// A constant declaration may specify a type as well as
const noDelay time.Duration = 0

// But in the absence of an explicit type,
// the type is inferred from the expression
// on the right hand side.
const timeout = 5 * time.Minute

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
}
