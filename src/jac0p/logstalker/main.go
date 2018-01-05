package main

import (
    "os"
    "fmt"
    "github.com/hpcloud/tail"
)

func main() {
    ch := make(chan string) // channel initialization
    f, _ := os.Create("/Users/jac0p/_Kompi/GO/src/jac0p/logstalker/newlog.log") // target log file
    for _, sf := range os.Args[1:] {
        go stalkFile(sf, ch)
    }

    // loop to read new lines received. Proposition for future functionality below.
    for {
        // fmt.Println(<-ch) // prints lines received from channel
        f.WriteString(<-ch + "\n")
    }
}

func stalkFile(source string, ch chan<- string) {
    t, err := tail.TailFile(source, tail.Config{Follow: true}) // tail source file
    if err != nil {
        ch <- fmt.Sprintf("error while reading %s: %v", source, err) // send error to channel
    }
    for line := range t.Lines {
        ch <- fmt.Sprintf("%s | %s", source, line)
    }
}
