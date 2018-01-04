package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    start := time.Now()
    ch := make(chan string)
    f, err := os.Create("/Users/jac0p/_Kompi/GO/src/gopl.io/ch1/fetchall/out")
    check(err)

    for _, url := range os.Args[1:] {
        if strings.HasPrefix(url, "https://") == false {
            url = "https://" + url
        }
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        // fmt.Println(<-ch) // receive from channel ch
        f.WriteString(<-ch + "\n")
    }
    // fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
    f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
    f.Close()
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
