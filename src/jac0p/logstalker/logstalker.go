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

    // goroutine to start reading files
    for _, src := range sources {
        if CheckIfFile(src) {
            fmt.Printf("YEAAAH: %s\n", src)
            // go stalkFile(src, ch)
        }
    }
}

func ChkParams(args []string) ([]string) {
    if len(args) < 1 {
        fmt.Println("application execution: ./logstalker -out=(log|out|all) log_file1 log_file2 log_file3")
        os.Exit(1)
    }
    return args
}

func CreateTG() {
    path, _ := filepath.Abs("./data")
    os.Create(path + "/newlog.log") // hardcoded for now
    // f, _ := os.Create(path + "/newlog.log") 
}

func CheckIfFile(src string) bool {
    f, err := os.Stat(src)
    if err != nil {
        fmt.Printf("error running stat on %s\n", src)
        return false
    }
    switch fm := f.Mode(); {
    case fm.IsDir():
        return false
    case fm.IsRegular():
        return true
    }
    panic("this should never happen")
}

func WalkDir(dir string) []string {
    o := []string{}
    filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
        if CheckIfFile(path) {
            o = append(o, path)
        }
        return nil
    })
    return o
}

func stalkFile(src string, ch chan<- string) {
    // tails file and sends results to channel
    t, err := tail.TailFile(src, tail.Config{Follow: true}) // tail source file
    if err != nil {
        ch <- fmt.Sprintf("error while reading %s: %v", src, err) // send error to channel
    }
    for line := range t.Lines {
        ch <- fmt.Sprintf("%s | %s", src, line)
    }
}

func readCH() {
    // loop to read new lines received
    for {
        // fmt.Println(<-ch) // prints lines received from channel
        // f.WriteString(<-ch + "\n")
    }
}
