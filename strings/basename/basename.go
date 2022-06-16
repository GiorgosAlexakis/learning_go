package main

import (
	"fmt"
	"strings"
)

// basename removes directory components and a .suffix.
// eg.,  a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename1(s string) string {
	// discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// like basename but uses the strings.LastIndex library function
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename1("/a/b/c.go"))
	fmt.Println(basename2("/a/b/c.go"))
}
