package main

import (
    "fmt"
    "path/filepath"
    "os"
)

func main() {
    path, _ := filepath.Abs("../logstalker/data")
    fmt.Println(path)
    err := os.Remove(path + "newlog.log")
    if err != nil {
        fmt.Print
    }
}
