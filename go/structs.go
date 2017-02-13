package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Id        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerId int
}

var dilbert Employee

func main() {
	dilbert.Id = 42

	fmt.Println(dilbert.Id)
}
