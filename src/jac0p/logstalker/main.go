package main

import (
  "fmt"
  "github.com/hpcloud/tail"
)

func main() {
  t, _ := tail.TailFile("/Users/jac0p/_Kompi/GO/src/jac0p/logstalker/mylog.log", tail.Config{Follow: true})
  for line := range t.Lines {
    fmt.Println(line.Text)
  }
}
