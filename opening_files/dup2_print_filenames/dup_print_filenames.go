// Dup_print_filenames prints the count and text of lines that appear more than once
// in the input and also prints the names of files in which such lines exist
// It reads from a list of named files
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	fmt.Println(filenames)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup_print_filenames: %v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
