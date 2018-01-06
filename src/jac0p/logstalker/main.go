package main

import (
    "os"
    "fmt"
    "flag"
    "github.com/hpcloud/tail"
)

func main() {
    // ch := make(chan string) // channel initialization

    sources, target := chkParams()
    fmt.Printf("SOURCES: %s", sources)
    fmt.Printf("TARGET: %s", target)

    // goroutine here to start reading files => startStalk function contains code

}


// checks for 'out' flag and returns source files & target logging method based value
func chkParams() ([]string, string) {
    outPtr := flag.String("out", "out", "indicates where the application should send it's output")
    if len(os.Args) < 2 {
        fmt.Println("application execution: ./logstalker -out=(log|out|all) log_file1 log_file2 log_file3")
        os.Exit(1)
    } else {
        if outPtr == nil {
            sources := os.Args[1:]
            return sources, "out"
        } else {
            sources := os.Args[2:]
            return sources, *outPtr
        }
    }
    panic("this should never happen")
}


func crtTG() {
    // f, _ := os.Create("/Users/jac0p/_Kompi/GO/src/jac0p/logstalker/newlog.log") // hardcoded for now
}

func startStalk(ch chan<- string) {
    // loops through cli params and starts tailing concurrently
    for _, sf := range os.Args[1:] {
        go stalkFile(sf, ch)
    }
}

func stalkFile(source string, ch chan<- string) {
    // tails file and sends results to channel
    t, err := tail.TailFile(source, tail.Config{Follow: true}) // tail source file
    if err != nil {
        ch <- fmt.Sprintf("error while reading %s: %v", source, err) // send error to channel
    }
    for line := range t.Lines {
        ch <- fmt.Sprintf("%s | %s", source, line)
    }
}

func readCH() {
    // loop to read new lines received
    for {
        // fmt.Println(<-ch) // prints lines received from channel
        // f.WriteString(<-ch + "\n")
    }
}
