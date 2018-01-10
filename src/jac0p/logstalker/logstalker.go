package logstalker

import (
    "io"
    "os"
    "fmt"
    "flag"
    "github.com/hpcloud/tail"
    "path/filepath"
)

const target = "/aggregated.log"

func Run() {
    params, w := chkParams()
    sources := findSources(params)
    ch := make(chan string) // channel initialization

    for _, src := range sources {
        go stalkFile(src, ch)
    }

    for {
        readCH(w, ch)
    }
}

func chkParams() ([]string, io.Writer) {
    outPtr := flag.String("out", "out", "indicates where the application should send it's output")
    flag.Parse() // parse flags || TODO: will need to return this 
    args := flag.Args()
    var w io.Writer

    if len(args) < 1 {
        fmt.Println("application execution: ./logstalker -out=(log|out|all) log_file1 log_file2 log_file3")
        os.Exit(1)
    }

    switch *outPtr {
    case "log":
        fmt.Printf("sending output to %s\n", target)
        w := createTG()
        return args, w
    case "out":
        fmt.Println("sending output to STDOUT")
        w := os.Stdout
        return args, w
    case "all":
        fmt.Printf("sending output to STDOUT and %s\n", target)
        w := io.MultiWriter(createTG(), os.Stdout)
        return args, w
    default:
        w := os.Stdout
        return args, w
    }
    return args, w
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

func createTG() (*os.File) {
    path, _ := filepath.Abs("./data")
    tg, err := os.Create(path + target) // hardcoded for now
    if err != nil {
        fmt.Errorf("unable to create %s: %s", target, err)
        os.Exit(1)
    }
    return tg
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

func readCH(w io.Writer, ch <-chan string) {
    // loop to read new lines received from channel
    for {
        // fmt.Println(<-ch) // prints lines
        fmt.Fprintf(w, "%s\n", <-ch)
    }
}
