package main

import (
    "fmt"
    // "os"
    "jac0p/logstalker"
    "jac0p/tokenizer"
)


func main() {
    logstalker.Run()
}

func testTokenizer() {
    const sentence = "I am a new string full of new opportunities"

    iter := tokenizer.Split(sentence, " ")

    for {
        token, ok := iter()
        if !ok {
            break
        }
        fmt.Println(token)
    }
}
