package main

import "fmt"

func main() {
    // Integers
    var b int8 = 127
    var c int16 = 32767
    var d int32 = 2147483647
    var e int64 = 9223372036854775807

    // Floats
    var f float32 = 3.14
    var g float64 = 3.141592653589793

    // Boolean
    var flag bool = true

    // String
    var message string = "Hello, Go!"

    // Print all variables in a single statement
    fmt.Printf("b: %d, c: %d, d: %d, e: %d, f: %f, g: %f, flag: %t, message: %s\n", b, c, d, e, f, g, flag, message)
}
