package main

import (
	"bytes"
	"fmt"
)

// comma inserts comma in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// non-recursive comma function, using bytes.Buffer instead of string concatenation
func commaWithBuffer(s string) string {
	n := len(s)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

// extending commaWithBuffer to handle floating point numbers
func commaWithFloatingNumbers(s string) string {
	n := len(s)
	var buf bytes.Buffer
	ind := len(s)
	neg := 0
	for i := 0; i < n; i++ {
		if s[i] == byte('-') {
			neg = 1
			continue
		}
		if s[i] == byte('.') {
			ind = i
		}
		if (n-i+neg)%3 == 0 && i < ind && i != neg {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	if neg == 1 {
		return "-" + buf.String()
	} else {
		return buf.String()
	}
}

func main() {
	s := "1234512311231312"
	fmt.Println(comma(s))
	fmt.Println(commaWithBuffer(s))
	fmt.Println(commaWithFloatingNumbers(s))
	s = "123451231231312.1241412"
	fmt.Println(commaWithFloatingNumbers(s))
	s = "-123451231231312.1241412"
	fmt.Println(commaWithFloatingNumbers(s))
}
