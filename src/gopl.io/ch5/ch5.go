package ch5


import (
    "fmt"
    "sort"
    "os"
)

func Sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

func Min(vals ...int) int {
    if len(vals) < 1 {
        fmt.Errorf("Min(%q) - Min requires at least 1 argument to run", vals)
        os.Exit(1)
    }
    sort.Ints(vals)
    return vals[0]
}

func Max(vals ...int) int {
    if len(vals) < 1 {
        fmt.Errorf("Max(%q) - Max requires at least 1 argument to run", vals)
        os.Exit(1)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(vals)))
    return vals[0]
}
