/*
Write a function CountingWriter with the signature below that, given an
io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable
that at any moment contains the number of bytes written to the new Writer.
*/
package main

import (
	"bytes"
	"fmt"
	"io"
)

type CounterWriter struct {
	writer io.Writer
	count  int64
}

func (w *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{
		writer: w,
		count:  0,
	}
	return cw, &cw.count
}

func main() {
	// An artificial input source.
	const input = "Now is the winter of discontent,\nMade glorious summer by this sun of York.\n"
	// use an io.Writer to count the number of bytes written to it.
	var buf bytes.Buffer
	// write to buf using a CounterWriter
	newWriter, count := CountingWriter(&buf)
	newWriter.Write([]byte(input))
	fmt.Println(*count)
}
