package main

import (
    "fmt"
    "github.com/hpcloud/tail"
)

func main() {
    // TODO: read from multiple source logs by calling 'tailFile' in a goroutine. The files
    // should be parsed from the command line. Target file should be hardcoded for now. 

    // vars
    ch := make(chan string) // channel
    // t, _ := tail.TailFile("/Users/jac0p/_Kompi/GO/src/jac0p/logstalker/mylog.log", tail.Config{Follow: true}) // tail source file

    // pseudo goroutine
    for _, sourceFile := range os.Args[1:] {
        go tailFile(sourceFile) // TODO: shorten variable name
    }


    f, err := os.Create("/Users/jac0p/_Kompi/GO/src/jac0p/logstalker/newlog.log") // target log file


    // loop to read new lines received. Proposition for future functionality below.
    for line := range t.Lines {
        fmt.Println(line.Text)
        // fmt.Println(<-ch) // TODO: receive lines from channel either in this loop or another 
        // f.WriteString(<-ch + "\n") TODO: eventually write results to target log
    }
}

func tailFile(file string, ch chan<- string) {
    // do stuff here
}
