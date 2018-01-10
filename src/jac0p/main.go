package main

import (
    "fmt"
    "jac0p/tokenizer"
)


func main() {
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
