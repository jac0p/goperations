package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)

    fmt.Print("ENTER INPUT: ")
    // text := input.Scan()
    for input.Scan() {
        counts[input.Text()]++
    }

    // fmt.Println(text)
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
