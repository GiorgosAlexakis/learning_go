# dup1
## **bufio package**

Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

The line: 

    input := bufio.NewScanner(os.Stdin)

Reads from the standard input stream with the use of a "Scanner" from the bufio package. This line also uses a buffer :

    type Reader struct {
    buf          []byte
	rd           io.Reader // reader provided by the client
    ...
}

It has a default size of 4096 bytes:

    const (
	defaultBufSize = 4096
    )

Next we use the Scan method which advances the Scanner to the next token, which will then be available through the Bytes or Text method. The Scan function returns true if there is a line and false when there is no more input.

    for input.Scan() {
        counts[input.Text()]++
    }

So basically for every line read (we use .Text() in order to retrieve the result) we check if the line's text has already been seen before(with map). Each call to input.Scan() also removes the newline character from the end.

Then we just print each line's text that appeared more than once in the standard input:

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }

### Printf and Println :
A note regarding the difference of printf to println (or even regarding functions such as log.Printf and fmt.Errorf) is that the first type formats the output in C style (%s for strings, %d for integers etc) . The second type does not format the output and leaves it as is. 
