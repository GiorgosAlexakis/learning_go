## **bufio package**

Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

The line: 

    input := bufio.NewScanner(os.Stdin)

Reads from the standard input stream with the use of a "Scanner" from the bufio package.