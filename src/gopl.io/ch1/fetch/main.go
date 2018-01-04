package main

import (
    "fmt"
    "io"
    // "io/ioutil"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if strings.HasPrefix(url, "http://") == false {
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        s := resp.Status
        // b, err := ioutil.ReadAll(resp.Body)
        b, err := io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", b)
        fmt.Printf("%s", s)

    }
}
