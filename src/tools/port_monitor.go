package main

import (
  "fmt"
  "os"
  "log"
  "os/exec"
)

func main() {
    log.Println("checking for open ports")
    cmd := "/usr/sbin/netstat"
    args := []string{"-ap", "tcp"}

    //out, err := exec.Command("/usr/sbin/netstat -ap tcp").Output()

    if err := exec.Command(cmd, args...).Run(); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
    }
    fmt.Println("Success!")

    out, err := exec.Command(cmd, args...).Output()
    if err != nil { log.Fatal(err) }
    log.Println("the following TCP ports are listening")
    log.Println(string(out))
}
