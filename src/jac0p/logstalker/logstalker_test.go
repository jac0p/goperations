package logstalker_test

import (
    // "fmt"
    "os"
    "testing"
    "jac0p/logstalker"
    "path/filepath"
    // "sort"
)

// ch := make(chan string) // global channel 


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
    tg := path + "/aggregated.log"
    os.Remove(tg) // remove existing file before testing
    logstalker.CreateTG() // create new file
    if _, err := os.Stat(tg); os.IsNotExist(err) {
        t.Errorf("CreateTG() didn't create: %s", tg)
    }
}

func TestCheckIfFile(t *testing.T) {
    var tests = []struct {
        input string
        want bool
    }{
        {"/tmp/testing/afile", true},
        {"/tmp/testing/bfile", true},
        {"/tmp/testing/adir", false},
    }

    for _, test := range tests {
        if logstalker.CheckIfFile(test.input) != test.want {
            t.Errorf("CheckIfFile(%s) != %s", test.input, test.want)
        }
    }
}

func TestWalkDir(t *testing.T) {
    input := "/tmp/testing/"
    want := []string{"/tmp/testing/afile", "/tmp/testing/bfile", "/tmp/testing/adir/cfile", "/tmp/testing/adir/bdir/dfile"}
    fl := logstalker.WalkDir(input)
    // sort.Strings(fl)

    // this test doesn't really tell all the truth
    if len(fl) != len(want) {
        t.Errorf("WalkDir(%q) = %v", input, fl)
    }
}




