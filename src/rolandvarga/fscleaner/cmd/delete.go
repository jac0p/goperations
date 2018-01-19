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

package cmd

import (
    // "fmt"
    "os"
    "sort"

    oshelper "jac0p/helper/os"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Deletes the provided folder and all of it's contents",
    Run: func(cmd *cobra.Command, args []string) {
        Run()
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}


func Run() {
    if !oshelper.CheckIfDir(srcDir) {
        log.Error("provided resource is not a directory or you have no permission to view it")
        os.Exit(1)
    }

    chdObjects := getChildObjects(srcDir) // get sorted list of directory objects
    delObjects := getToDelete(chdObjects) // get list of doomed objects
    remove(delObjects)
}


func getChildObjects(src string) []string {
    return sortObjects(oshelper.ListDir(srcDir))
}

func sortObjects(obj []string) []string {
    if rvsList {
        // deletes new elements first
        sort.Sort(sort.Reverse(sort.StringSlice(obj)))
    } else {
        // deletes old elements first
        sort.Strings(obj)
    }
    return obj
}

func getToDelete(chdObjects []string) []string {
    chdCount := len(chdObjects)
    keepObjects := chdObjects[chdCount-keepCnt:]

    for i := 0; i < len(chdObjects); i++ {
        for _, k := range keepObjects {
            if chdObjects[i] == k {
                chdObjects = append(chdObjects[:i], chdObjects[i+1:]...)
                i--      // decrease index so we don't skip any elements
            }
        }
    }
    return chdObjects
}

func remove(list []string) {
    if hrdDel {
        for _, elem := range list {
            s := srcDir + "/" + elem
            log.Info("deleting: " + s)
            oshelper.DeleteDir(s)
        }
    }
    if sftDel {
        // something like
        // oshelper.MoveDir(s, ".Trash/s")
    }
}

