package logstalker

import (
    "os"
    "fmt"
    "flag"
    "github.com/hpcloud/tail"
    "path/filepath"
)


var outPtr = flag.String("out", "out", "indicates where the application should send it's output")

func Start() {
    // ch := make(chan string) // channel initialization
    flag.Parse() // parse flags

    sources := ChkParams(flag.Args()) // sends arguments without flags
    fmt.Printf("SOURCES: %s\n", sources)
    fmt.Printf("FLAG: %v\n", *outPtr)

    // goroutine here to start reading files => startStalk function contains code
}

func ChkParams(args []string) ([]string) {
    if len(args) < 1 {
        fmt.Println("application execution: ./logstalker -out=(log|out|all) log_file1 log_file2 log_file3")
        os.Exit(1)
    }
    return args
}

func CreateTG() {
    // f, _ := os.Create(filepath.Abs("./data") + "/newlog.log") 
    path, _ := filepath.Abs("./data")
    os.Create(path + "/newlog.log") // hardcoded for now
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
