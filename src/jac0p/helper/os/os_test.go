package os_test

import (
    // "fmt"
    "os"
    "testing"
    // "path/filepath"
    // "sort"

    oshelper "jac0p/helper/os"
    log "github.com/sirupsen/logrus"
)

func TestCreateTG(t *testing.T) {
    // path, _ := filepath.Abs("./data")
    // tg := path + "/aggregated.log"
    tg := "/tmp/testfile.log"
    os.Remove(tg) // remove existing file before testing
    oshelper.CreateTG(tg) // create new file
    if _, err := os.Stat(tg); os.IsNotExist(err) {
        t.Errorf("CreateTG() didn't create: %s", tg)
    }
}

func TestCheckIfFile(t *testing.T) {
    var tests = []struct {
        input string
        want bool
    }{
        {"./cases/afile", true},
        {"./cases/bfile", true},
        {"./cases/adir", false},
    }

    for _, test := range tests {
        if oshelper.CheckIfFile(test.input) != test.want {
            t.Errorf("CheckIfFile(%s) != %s", test.input, test.want)
        }
    }
}

func TestWalkDir(t *testing.T) {
    input := "./cases/walkme/"
    want := []string{"./cases/walkme/afile", "./cases/walkme/bfile", "./cases/walkme/adir/cfile", "./cases/walkme/adir/bdir/dfile"}
    fl := oshelper.WalkDir(input)
    // sort.Strings(fl)

    // this test doesn't really tell all the truth
    if len(fl) != len(want) {
        t.Errorf("WalkDir(%q) = %v", input, fl)
    }
}

func TestDeleteDir(t *testing.T) {
    target := "/tmp/deleteme/"
    log.Debug("creating directory: %s", target)
    err := os.MkdirAll(target, 0666)
    if err != nil {
        t.Errorf("cannot create %s", target)
    }

    oshelper.DeleteDir("/tmp/deleteme/")
    if oshelper.CheckIfDir(target) {
        t.Errorf("directory wasn't deleted..")
    }

}





