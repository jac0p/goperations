package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    start := time.Now()

    // var s, sep string
    for i := 1; i < len(os.Args); i++ {
        fmt.Println("INDEX: ", i, " ARG: ", os.Args[i])
    }

    for i, arg := range os.Args[1:] {
        fmt.Println("INDEX: ", i, " ARG: ", arg)
    }

    secs := time.Since(start).Seconds()
    fmt.Printf("TIME ELAPSED: %.2fs", secs)
}
