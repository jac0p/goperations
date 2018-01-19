// Copyright © 2018 Roland Varga <roland.varga.can@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package os

import (
	// "fmt"
  "os"
  "io/ioutil"

  "path/filepath"
  log "github.com/sirupsen/logrus"
)

// checks if src is a file
func CheckIfFile(src string) bool {
    f, err := os.Stat(src)
    if err != nil {
        log.Errorf("error running stat on %s", src)
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

// checks if src is a directory
func CheckIfDir(src string) bool {
    f, err := os.Stat(src)
    if err != nil {
        log.Errorf("error running stat on %s", src)
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

// walks provided directory and collects all files recursively
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

// creates a file in absolute path
func CreateTG(tgt string) (*os.File) {
    // path, _ := filepath.Abs("./data")
    tg, err := os.Create(tgt) // hardcoded for now
    if err != nil {
        log.Errorf("unable to create %s: %s", tgt, err)
        os.Exit(1)
    }
    return tg
}

func ListDir(src string) []string {
    o := []string{}
    log.Infof("listing contents of %s", src)
    cnt, err := ioutil.ReadDir(src)
    if err != nil {
        log.Errorf("unable to list %s", src)
    }
    for _, c := range cnt {
        o = append(o, c.Name())
    }
    return o
}

func DeleteDir(tgt string) {
    log.Infof("attempting to delete %s", tgt)
    err := os.RemoveAll(tgt)
    if err != nil {
        log.Errorf("unable to delete %s", tgt)
    }
    log.Info("directory deleted.")
}




