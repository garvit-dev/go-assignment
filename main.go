package main

import (
	employee "Assignment/employee"
	// employee "github.com/garvit-dev/go-assignment/tree/master/employee"

	aqua "github.com/rightjoin/fuel"
)

func main() {

	server := aqua.NewServer()
	server.AddService(&employee.EmployeeService{})
	server.Run()
}
