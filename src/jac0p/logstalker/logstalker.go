package logstalker

import (
    "os"
    "fmt"
    "flag"
    "github.com/hpcloud/tail"
    "path/filepath"
)

const target = "/aggregated.log"

func Run() {
    params := chkParams()
    sources := findSources(params)

    // create target
    fmt.Printf("creating target file: %s\n", target)
    createTG()

    ch := make(chan string) // channel initialization
    for _, src := range sources {
        go stalkFile(src, ch)
    }

    // forward to target
    for {
        fmt.Println(<-ch) // prints lines received from channel
    }
}

func chkParams() ([]string) {
    flag.String("out", "out", "indicates where the application should send it's output")
    flag.Parse() // parse flags || TODO: will need to return this 
    args := flag.Args()

    if len(args) < 1 {
        fmt.Println("application execution: ./logstalker -out=(log|out|all) log_file1 log_file2 log_file3")
        os.Exit(1)
    }
    return args
}

func findSources(params []string) ([]string) {
    sources := []string{}
    for _, p := range params {
        if checkIfFile(p) {
            sources = append(sources, p)
        }
        if checkIfDir(p) {
            sources = append(sources, walkDir(p)...)
        }
    }
    return sources
}

// SOOOO DRYYYYYY
func checkIfFile(src string) bool {
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


// SOOOO DRYYYYYY
func checkIfDir(src string) bool {
    f, err := os.Stat(src)
    if err != nil {
        fmt.Printf("error running stat on %s\n", src)
        return false
    }
    switch fm := f.Mode(); {
    case fm.IsDir():
        return true
    case fm.IsRegular():
        return false
    }
    panic("this should never happen")
}

func walkDir(dir string) []string {
    o := []string{}
    filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
        if checkIfFile(path) {
            o = append(o, path)
        }
        return nil
    })
    return o
}

func createTG() {
    path, _ := filepath.Abs("./data")
    os.Create(path + target) // hardcoded for now
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

// func readCH(ch <-chan string) {
//     // loop to read new lines received
//     for {
//         fmt.Println(<-ch) // prints lines received from channel
//         // f.WriteString(<-ch + "\n")
//     }
// }
