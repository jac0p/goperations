package ch4

import (
    "fmt"
    "time"
)

type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

func Run() {
    var dilbert Employee
    fmt.Println(dilbert)
    dilbert.Salary += 75000
    fmt.Println(dilbert)
    dilbert.Salary -= 5000
    fmt.Println(dilbert)

    position := &dilbert.Position
    *position = "Architect"
    *position = "Senior " + *position
    fmt.Println(dilbert)

    var employeeOfTheMonth *Employee = &dilbert
    employeeOfTheMonth.Position += " (pro)"
    fmt.Println(dilbert)

}

