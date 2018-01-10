package tokenizer

import "strings"

type tokenizer func() (token string, ok bool)

func Split(s, sep string) tokenizer {
    tokens, last := strings.Split(s, sep), 0

    return func() (string, bool) {
        if len(tokens) == last {
            return "", false
        }
        last = last + 1
        return tokens[last-1], true
    }
}
