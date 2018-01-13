package main

import (
    "fmt"
    "gopl.io/ch4"
    // "gopl.io/ch5"
    "os"
)


func f(...int) { }
func g([]int) { }

func errorf(linenum int, format string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
    fmt.Fprintf(os.Stderr, format, args...)
    fmt.Fprintln(os.Stderr)
}

func myJoin(args ...string) string {
    str := ""
    for _, a := range args {
        str = str + string(a)
    }
    return str
}

func main() {
    fmt.Println("========================================")
    ch4.Run()

}
