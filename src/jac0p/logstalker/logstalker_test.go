package logstalker_test

import (
    // "fmt"
    "os"
    "testing"
    "jac0p/logstalker"
    "path/filepath"
)


func TestSourceMatchChkParams(t *testing.T) {
    var tests = []struct {
        input []string
        sources []string
    }{
        {[]string{"log1.log"}, []string{"log1.log"}},
        {[]string{"log1.log", "log2.log"}, []string{"log1.log", "log2.log"}},
    }

    for _, test := range tests {
        s := logstalker.ChkParams(test.input)

        // checks if number of sources match
        if len(s) != len(test.sources) {
            t.Errorf("ChkParams(%q) = %v", test.input, s)
        }

        // check if sources are identical
        for i := range s {
            if s[i] != test.sources[i] {
                t.Errorf("ChkParams(%q) = %v", test.input, s)
            }
        }
    }
}

func TestCreateTG(t *testing.T) {
    path, _ := filepath.Abs("./data")
    tg := path + "/newlog.log"
    os.Remove(tg) // remove existing file before testing
    logstalker.CreateTG() // create new file
    if _, err := os.Stat(tg); os.IsNotExist(err) {
        t.Errorf("CreateTG() didn't create: %s", tg)
    }



}
