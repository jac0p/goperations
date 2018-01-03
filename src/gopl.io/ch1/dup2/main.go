package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]

    for _, arg := range files {
        f, err := os.Open(arg)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
            continue
        }
        countLines(f, counts, arg)
        f.Close()
    }
}

func countLines(f *os.File, counts map[string]int, arg string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%s\t%s\t%d\n", arg, line, n)
        }
    }
}
