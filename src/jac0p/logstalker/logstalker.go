package logstalker

import (
    "os"
    "fmt"
    "flag"
    "github.com/hpcloud/tail"
    "path/filepath"
)


func Start() {
    // ch := make(chan string) // channel initialization
    flag.String("out", "out", "indicates where the application should send it's output")
    flag.Parse() // parse flags
    params := ChkParams(flag.Args()) // sends arguments without flags
    sources := []string{}

    // find all source files
    for _, p := range params {
        if CheckIfFile(p) {
            sources = append(sources, p)
        }
        if CheckIfDir(p) {
            sources = append(sources, WalkDir(p)...)
        }
    }

    // go StalkFile(sources, ch)
}

func ChkParams(args []string) ([]string) {
    if len(args) < 1 {
        fmt.Println("application execution: ./logstalker -out=(log|out|all) log_file1 log_file2 log_file3")
        os.Exit(1)
    }
    return args
}

// SOOOO DRYYYYYY
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

// SOOOO DRYYYYYY
func CheckIfDir(src string) bool {
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

func CreateTG() {
    path, _ := filepath.Abs("./data")
    os.Create(path + "/aggregated.log") // hardcoded for now
    // f, _ := os.Create(path + "/newlog.log") 
}

func StalkFile(src string, ch chan<- string) {
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
