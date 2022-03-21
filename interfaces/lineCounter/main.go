package main

import (
	"bufio"
	"fmt"
	"strings"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c += LineCounter(1)
	}
	return int(*c), nil
}

func main() {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	var c LineCounter
	c.Write([]byte(input))
	fmt.Println(c)
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
