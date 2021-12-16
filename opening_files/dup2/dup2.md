# dup2

dup2 serves the same function as dup1 but has the added functionality of being able to read from a list of named files.

#### **os.Open()**
    f, err := os.Open(arg)

The first value is an open file (*os.File) that is used in subsequent reads by the Scanner. The second result of os.Open() is the build-in error type. To check if the file was opened succesfully we need to compare it with the value 'nil' . If the second result is equal to nil then we have succesfully opened the file without any problems.

Otherwise, if the err is not equal to 'nil' , the error value describes the problem. In this case we handle the error by printing a message on the standard error stream using Fprintf and the verb verb %v which displays a value of any type with the default format. The program then carries on with the next file; the continue statement  goes to the next iteration of the enclosing for loop.

