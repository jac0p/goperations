package main

import (
    "fmt"
    "gopl.io/ch5"
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
    fmt.Println(ch5.Sum())
    fmt.Println(ch5.Sum(3))
    fmt.Println(ch5.Sum(1,2, 3, 4))

    values := []int{5, 6, 7, 8}
    fmt.Println(ch5.Sum(values...))

    fmt.Printf("%T\n", f)
    fmt.Printf("%T\n", g)
    fmt.Printf("%T\n", values)

    fmt.Println("========================================")

    linenum, name := 12, "count"
    errorf(linenum, "undefined: %s", name)

    fmt.Println("========================================")
    fmt.Println(myJoin("a", "b"))

    fmt.Println("========================================")
    fmt.Println(ch5.Min(1))
    fmt.Println(ch5.Min(1, 2, 3, 4))
    fmt.Println(ch5.Min(5, 2, 3, 4))

    fmt.Println(ch5.Max(1, 2))
    fmt.Println(ch5.Max(1, 2, 3, 4))
    fmt.Println(ch5.Max(5, 2, 3, 4))

}
