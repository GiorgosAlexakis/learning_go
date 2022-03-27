/*
The LimitReader function in the io package accepts an io.Reader r and a
number of bytes n, and returns another Reader that reads from r but reports an end-of-file
condition after n bytes
*/
package main

import "io"

type limitedReader struct {
	r io.Reader
	n int64
}

func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[0:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{r, n}
}
